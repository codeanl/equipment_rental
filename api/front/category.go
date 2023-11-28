package front

import (
	"github.com/gin-gonic/gin"
	"outdoor_rental/api"
	"outdoor_rental/model/req"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type Category struct {
}

// GetList 列表
func (*Category) GetList(c *gin.Context) {
	r.SuccessData(c, api.CategoryService.GetList(utils.BindQuery[req.FrontCategoryList](c)))
}

// GetList 列表
func (*Category) GetListNextCateAndSpu(c *gin.Context) {
	r.SuccessData(c, api.CategoryService.GetListNextCateAndSpu(utils.BindQuery[req.FrontCategoryList](c)))
}
