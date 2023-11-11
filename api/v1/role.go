package v1

import (
	"github.com/gin-gonic/gin"
	"outdoor_rental/model/req"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type Role struct{}

// SaveOrUpdate 添加｜｜更新菜单
func (*Role) SaveOrUpdate(c *gin.Context) {
	r.SendCode(c, roleService.SaveOrUpdate(utils.BindValidJson[req.SaveOrUpdateRole](c)))
}

// Delete 删除
func (*Role) Delete(c *gin.Context) {
	r.SendCode(c, roleService.Delete(utils.BindJson[req.Delete](c)))
}

// GetTreeList 列表
func (*Role) GetTreeList(c *gin.Context) {
	r.SuccessData(c, roleService.GetTreeList(utils.BindQuery[req.RoleList](c)))
}
