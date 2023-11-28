package front

import (
	"github.com/gin-gonic/gin"
	"outdoor_rental/api"
	"outdoor_rental/model/req"
	"outdoor_rental/utils/r"
)

type Slideshow struct {
}

//SlideshowList 列表
func (*Slideshow) SlideshowList(c *gin.Context) {
	r.SuccessData(c, api.SlideShowService.SlideshowList(req.SlideshowList{Status: "1"}))
}
