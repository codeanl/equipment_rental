package front

import (
	"github.com/gin-gonic/gin"
	"outdoor_rental/api"
	"outdoor_rental/model/req"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type Article struct {
}

// SaveOrUpdateCart 添加｜｜更新菜单
func (*Article) SaveOrUpdateArticle(c *gin.Context) {
	r.SendCode(c, api.ArticleService.SaveOrUpdateArticle(utils.BindJson[req.SaveOrUpdateArticle](c)))
}

// DeleteCart 删除
func (*Article) DeleteArticle(c *gin.Context) {
	r.SendCode(c, api.ArticleService.DeleteArticle(utils.BindJson[req.Delete](c)))
}

// GetCartList 列表
func (*Article) GetArticleList(c *gin.Context) {
	r.SuccessData(c, api.ArticleService.GetArticleList(utils.BindQuery[req.ArticleList](c)))
}
