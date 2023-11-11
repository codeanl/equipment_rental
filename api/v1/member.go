package v1

import (
	"github.com/gin-gonic/gin"
	"outdoor_rental/model/req"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
	"strconv"
)

type Member struct{}

// MemberInfo  获取会员信息
func (*Member) MemberInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	r.SuccessData(c, memberService.MemberInfo(id))
}

// MemberList 会员列表
func (*Member) MemberList(c *gin.Context) {
	r.SuccessData(c, memberService.MemberList(utils.BindQuery[req.MemberList](c)))
}

// MemberUpdate 更新会员
func (*Member) MemberUpdate(c *gin.Context) {
	r.SendCode(c, memberService.MemberUpdate(utils.BindJson[req.MemberUser](c)))
}

// MemberDelete 删除会员
func (*Member) MemberDelete(c *gin.Context) {
	r.SuccessData(c, memberService.MemberDelete(utils.BindJson[req.Delete](c)))
}
