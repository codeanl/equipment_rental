package v1

import (
	"github.com/gin-gonic/gin"
	"outdoor_rental/model/req"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type Menu struct{}

// SaveOrUpdate 添加｜｜更新菜单
func (*Menu) SaveOrUpdate(c *gin.Context) {
	r.SendCode(c, menuService.SaveOrUpdate(utils.BindJson[req.SaveOrUpdateMenu](c)))
}

// Delete 删除菜单
func (*Menu) Delete(c *gin.Context) {
	r.SendCode(c, menuService.Delete(utils.BindValidJson[req.Delete](c)))
}

// GetTreeList 菜单列表(树形)
func (*Menu) GetTreeList(c *gin.Context) {
	r.SuccessData(c, menuService.GetTreeList())
}

// 获取当前用户菜单: 生成后台管理界面的菜单
func (*Menu) GetUserMenu(c *gin.Context) {
	r.SuccessData(c, menuService.GetUserMenuTree(
		utils.GetFromContext[int](c, "id")))
}
