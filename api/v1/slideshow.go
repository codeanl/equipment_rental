package v1

import (
	"github.com/gin-gonic/gin"
	"outdoor_rental/api"
	"outdoor_rental/model/req"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type Slideshow struct{}

// SaveOrUpdate 添加｜｜更新
func (*Slideshow) SaveOrUpdate(c *gin.Context) {
	r.SendCode(c, api.SlideShowService.SaveOrUpdate(utils.BindJson[req.SaveOrUpdateSlideshow](c)))
}

// SlideshowDelete 删除
func (*Slideshow) SlideshowDelete(c *gin.Context) {
	r.SendCode(c, api.SlideShowService.Delete(utils.BindValidJson[req.Delete](c)))
}

//SlideshowList 列表
func (*Slideshow) SlideshowList(c *gin.Context) {
	r.SuccessData(c, api.SlideShowService.SlideshowList(utils.BindQuery[req.SlideshowList](c)))
}
