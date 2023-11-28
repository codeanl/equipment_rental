package v1

import (
	"github.com/gin-gonic/gin"
	"outdoor_rental/api"
	"outdoor_rental/model/req"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type Category struct{}

// SaveOrUpdate 添加｜｜更新菜单
func (*Category) SaveOrUpdate(c *gin.Context) {
	r.SendCode(c, api.CategoryService.SaveOrUpdate(utils.BindJson[req.SaveOrUpdateCategory](c)))
}

// Delete 删除
func (*Category) Delete(c *gin.Context) {
	r.SendCode(c, api.CategoryService.Delete(utils.BindJson[req.Delete](c)))
}

// GetTreeList 列表
func (*Category) GetTreeList(c *gin.Context) {
	r.SuccessData(c, api.CategoryService.GetTreeList())
}
