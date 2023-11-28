package front

import (
	"github.com/gin-gonic/gin"
	"outdoor_rental/api"
	"outdoor_rental/model/req"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type Member struct {
}

//会员登录

func (*Member) Login(c *gin.Context) {
	loginVo, code := api.MemberService.MemberLogin(c, utils.BindValidJson[req.MemberLogin](c))
	r.SendData(c, code, loginVo)
}
