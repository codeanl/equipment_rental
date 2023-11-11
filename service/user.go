package service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"outdoor_rental/config"
	"outdoor_rental/dao"
	"outdoor_rental/model"
	"outdoor_rental/model/dto"
	"outdoor_rental/model/req"
	"outdoor_rental/model/resp"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
	"time"
)

type User struct{}

// Login 登录
func (*User) Login(c *gin.Context, req req.Login) (token string, code int) {
	//查询用户是否存在
	userInfo := dao.GetOne(model.User{}, "username", req.Username)
	if userInfo.ID == 0 {
		return "", r.ERROR_USER_NOT_EXIST //该用户不存在
	}
	//校验密码
	if !utils.Encryptor.BcryptCheck(req.Password, userInfo.Password) {
		return "", r.ERROR_PASSWORD_WRONG //密码错误
	}
	//账号是否禁用
	if userInfo.Status != "1" {
		return "", r.ERROR_USER_DISABLE //用户已被锁定
	}
	// 获取用户详细信息
	roles := roleDao.GetRolesByUserId(userInfo.ID)
	//没有角色不可登录
	if len(roles) == 0 {
		return "", r.ERROR_USER_NO_ROLE //用户无角色，不可登录"
	}
	// 获取 IP 相关信息
	ipAddress := utils.IP.GetIpAddress(c)                 //ip
	ipSource := utils.IP.GetIpSourceSimpleIdle(ipAddress) //ip地址
	browser, os := "unknown", "unknown"                   //浏览器 操作系统
	if userAgent := utils.IP.GetUserAgent(c); userAgent != nil {
		browser = userAgent.Name + " " + userAgent.Version.String()
		os = userAgent.OS + " " + userAgent.OSVersion.String()
	}
	// 登录信息正确, 生成 Token
	// UUID 生成方法: ip + 浏览器信息 + 操作系统信息
	uuid := utils.Encryptor.MD5(ipAddress + browser + os)
	token, err := utils.GetJWT().GenToken(userInfo.ID, userInfo.Username, roles, uuid) //只设定一个角色, 获取第一个值
	if err != nil {
		utils.Logger.Info("登录时生成 Token 错误: ", zap.Error(err))
		return "", r.ERROR_TOKEN_CREATE
	}
	// 更新用户验证信息: ip 信息 + 上次登录时间
	dao.Update(&model.User{
		Public:        model.Public{ID: userInfo.ID},
		IpAddress:     ipAddress,
		IpSource:      ipSource,
		LastLoginTime: time.Now(),
	}, "ip_address", "ip_source", "last_login_time")
	//缓存信息
	redisInfo := dto.UserDetailDTO{
		LoginVO: resp.LoginVO{
			ID:            userInfo.ID,
			Email:         userInfo.Email,
			Username:      userInfo.Username,
			Nickname:      userInfo.Nickname,
			Avatar:        userInfo.Avatar,
			Intro:         userInfo.Intro,
			IpAddress:     ipAddress,
			IpSource:      ipSource,
			LastLoginTime: time.Now(),
		},
		Status:   userInfo.Status,
		Roles:    roles,
		Password: userInfo.Password,
		Browser:  browser,
		OS:       os,
	}
	redisInfo.Token = token
	// 保存用户信息到 Session 和 Redis 中
	sessionInfoStr := utils.Json.Marshal(redisInfo)
	utils.Redis.Set(KEY_USER+uuid, sessionInfoStr, time.Duration(config.Cfg.Session.MaxAge)*time.Second)
	//sessionInfoStr := utils.Json.Marshal(dto.SessionInfo{UserDetailDTO: userDetailDTO})
	//session := sessions.Default(c)
	//session.Set(KEY_USER+uuid, sessionInfoStr)
	//session.Save()
	return token, r.OK
}

//UserAdd 添加用户
func (*User) UserAdd(req req.UserAdd) (code int) {
	// 检查用户名已存在, 则该账号已经注册过
	userInfo := dao.GetOne(model.User{}, "username", req.Username)
	if userInfo.ID > 0 {
		return r.ERROR_USER_NAME_USED
	}
	user := &model.User{
		Username:      req.Username,
		Phone:         req.Phone,
		Password:      utils.Encryptor.BcryptHash(req.Password),
		Nickname:      "用户" + req.Username,
		Status:        "1",
		Avatar:        "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		LastLoginTime: time.Now(),
	}
	dao.Create(user)
	return r.OK
}

// UserInfo 用户详情
func (s *User) UserInfo(id int) resp.UserInfoVO {
	var userInfo model.User
	dao.GetOne(&userInfo, "id", id)
	data := utils.CopyProperties[resp.UserInfoVO](userInfo) //复制
	data.Roles = roleDao.GetRolesNameByUserId(id)
	//
	menuList := menuDao.GetMenusByUserInfoId(id)
	datamenu := utils.CopyProperties[[]resp.MenuListVO](menuList)
	data.Menus = s.buildMenuTree(datamenu, 0)
	return data
}

// UserUpdate 更新用户
func (*User) UserUpdate(req req.UpdateUser) int {
	userInfo := model.User{
		Public:   model.Public{ID: req.ID},
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Email:    req.Email,
		Phone:    req.Phone,
		Intro:    req.Intro,
		Status:   req.Status,
	}
	phone := dao.GetOne(model.User{}, "phone", req.Phone)
	if phone.ID > 0 && phone.ID != req.ID {
		return r.ERROR_PHONE_EXIST //该号码已绑定其他账户
	}
	dao.Update(&userInfo)
	// 清空 user_role 关系
	dao.Delete(model.UserRole{}, "user_id = ?", req.ID)
	if len(req.RoleIds) > 0 {
		// 要更新的 user_role 列表
		var userRoles []model.UserRole
		for _, id := range req.RoleIds {
			userRoles = append(userRoles, model.UserRole{
				RoleId: id,
				UserId: req.ID,
			})
		}
		dao.Create(&userRoles)
	}
	return r.OK
}

// UserDelete 删除用户
func (*User) UserDelete(req req.Delete) int {
	dao.Delete(&model.User{}, "id in (?)", req.ID)
	return r.OK
}

// UserList 用户列表
func (*User) UserList(req req.UserList) resp.PageResult[[]resp.UserListVO] {
	list, count, _ := userDao.UserList(req)
	return resp.PageResult[[]resp.UserListVO]{
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		Total:    count,
		List:     list,
	}
}

//UserSetPass 更新密码
func (*User) UserSetPass(id int, req req.SetPass) int {
	var userInfo model.User
	dao.GetOne(&userInfo, "id", id)
	if !utils.Encryptor.BcryptCheck(req.OldPass, userInfo.Password) {
		return r.ERROR_PASSWORD_WRONG
	}
	if req.NewPass != req.OnceNewPass {
		return r.ERROR_ONCE_PASSWORD
	}
	dao.Update(&model.User{
		Public:   model.Public{ID: id},
		Password: utils.Encryptor.BcryptHash(req.NewPass),
	}, "password")
	return r.OK
}

// UserUpdateInfo 更新个人信息
func (*User) UserUpdateInfo(id int, req req.UserUpdateInfo) int {
	userInfo := model.User{
		Public:   model.Public{ID: id},
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Email:    req.Email,
		Phone:    req.Phone,
		Intro:    req.Intro,
	}
	phone := dao.GetOne(model.User{}, "phone", req.Phone)
	if phone.ID > 0 && phone.ID != id {
		return r.ERROR_PHONE_EXIST //该号码已绑定其他账户
	}
	dao.Update(&userInfo)
	return r.OK
}

//
func (s *User) buildMenuTree(menuItems []resp.MenuListVO, parentID int) []resp.MenuListVO {
	tree := make([]resp.MenuListVO, 0)
	for _, item := range menuItems {
		if item.ParentId == parentID {
			children := s.buildMenuTree(menuItems, item.ID)
			item.Children = children
			tree = append(tree, item)
		}
	}
	return tree
}
