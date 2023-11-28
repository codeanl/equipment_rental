package service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"outdoor_rental/dao"
	"outdoor_rental/model"
	"outdoor_rental/model/req"
	"outdoor_rental/model/resp"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
	"time"
)

type Member struct{}

// MemberInfo 用户详情
func (s *Member) MemberInfo(id int) resp.MemberInfoVO {
	var memberInfo model.Member
	dao.GetOne(&memberInfo, "id", id)
	data := utils.CopyProperties[resp.MemberInfoVO](memberInfo) //复制
	return data
}

// MemberList 用户列表
func (*Member) MemberList(req req.MemberList) resp.PageResult[[]resp.MemberListVO] {
	list, count, _ := memberDao.MemberList(req)
	return resp.PageResult[[]resp.MemberListVO]{
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		Total:    count,
		List:     list,
	}
}

// MemberUpdate 更新用户
func (*Member) MemberUpdate(req req.MemberUser) int {
	userInfo := model.Member{
		Public:   model.Public{ID: req.ID},
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Email:    req.Email,
		Avatar:   req.Avatar,
		Intro:    req.Intro,
		Status:   req.Status,
	}
	dao.Update(&userInfo)
	return r.OK
}

// MemberDelete 删除用户
func (*Member) MemberDelete(req req.Delete) int {
	dao.Delete(&model.Member{}, "id in (?)", req.ID)
	return r.OK
}

//
//会员登录
func (*Member) MemberLogin(c *gin.Context, req req.MemberLogin) (respp resp.LoginMemberVO, code int) {
	//查询用户是否存在
	userInfo := dao.GetOne(model.Member{}, "phone", req.Phone)
	if userInfo.ID == 0 {
		return respp, r.ERROR_USER_NOT_EXIST //该用户不存在
	}
	//校验密码
	if !utils.Encryptor.BcryptCheck(req.Password, userInfo.Password) {
		return respp, r.ERROR_PASSWORD_WRONG //密码错误
	}
	//账号是否禁用
	if userInfo.Status != "1" {
		return respp, r.ERROR_USER_DISABLE //用户已被锁定
	}
	// 获取 IP 相关信息
	ipAddress := utils.IP.GetIpAddress(c)                 //ip
	ipSource := utils.IP.GetIpSourceSimpleIdle(ipAddress) //ip地址
	// 登录信息正确, 生成 Token
	// UUID 生成方法: ip + 浏览器信息 + 操作系统信息
	token, err := utils.GetJWT().GenTokenMember(userInfo.ID, userInfo.Phone) //只设定一个角色, 获取第一个值
	if err != nil {
		utils.Logger.Info("登录时生成 Token 错误: ", zap.Error(err))
		return respp, r.ERROR_TOKEN_CREATE
	}
	// 更新用户验证信息: ip 信息 + 上次登录时间
	dao.Update(&model.Member{
		Public:        model.Public{ID: userInfo.ID},
		IpAddress:     ipAddress,
		IpSource:      ipSource,
		LastLoginTime: time.Now(),
	}, "ip_address", "ip_source", "last_login_time")
	//
	data := utils.CopyProperties[resp.LoginMemberVO](userInfo)
	data.Token = token
	return data, r.OK
}

//会员注册

//更新个人信息
