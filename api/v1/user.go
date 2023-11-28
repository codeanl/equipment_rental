package v1

import (
	"github.com/gin-gonic/gin"
	"outdoor_rental/api"
	"outdoor_rental/model/req"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
	"strconv"
)

type User struct{}

// Upload 上传
func (*User) Upload(c *gin.Context) {
	code, url := utils.Upload(c)
	r.SendData(c, code, url)
}

// Login 登录
func (*User) Login(c *gin.Context) {
	loginVo, code := api.UserService.Login(c, utils.BindValidJson[req.Login](c))
	r.SendData(c, code, loginVo)
}

// Profile  获取个人信息
func (*User) Profile(c *gin.Context) {
	r.SuccessData(c, api.UserService.UserInfo(utils.GetFromContext[int](c, "id")))
}

// UserAdd 添加用户
func (*User) UserAdd(c *gin.Context) {
	r.SendCode(c, api.UserService.UserAdd(utils.BindValidJson[req.UserAdd](c)))
}

// UserUpdate 更新
func (*User) UserUpdate(c *gin.Context) {
	r.SendCode(c, api.UserService.UserUpdate(utils.BindJson[req.UpdateUser](c)))
}

// UserInfo 获取用户信息
func (*User) UserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	r.SuccessData(c, api.UserService.UserInfo(id))
}

// UserDelete 删除
func (*User) UserDelete(c *gin.Context) {
	r.SuccessData(c, api.UserService.UserDelete(utils.BindJson[req.Delete](c)))
}

// UserList 用户列表
func (*User) UserList(c *gin.Context) {
	r.SuccessData(c, api.UserService.UserList(utils.BindQuery[req.UserList](c)))
}

// UserSetPass 更新密码
func (*User) UserSetPass(c *gin.Context) {
	r.SendCode(c, api.UserService.UserSetPass(utils.GetFromContext[int](c, "id"), utils.BindValidJson[req.SetPass](c)))
}

// UserUpdateInfo 更新个人信息
func (*User) UserUpdateInfo(c *gin.Context) {
	r.SendCode(c, api.UserService.UserUpdateInfo(utils.GetFromContext[int](c, "id"), utils.BindValidJson[req.UserUpdateInfo](c)))
}
