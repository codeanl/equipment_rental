package r

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 响应结构体
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// 返回 JSON 数据
func ReturnJson(c *gin.Context, httpCode, code int, msg string, data any) {
	c.JSON(httpCode, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

// 自定义 httpCode, code, data
func Send(c *gin.Context, httpCode, code int, data any) {
	ReturnJson(c, httpCode, code, GetMsg(code), data)
}

// 自动根据 code 获取 message, 且 data == nil
func SendCode(c *gin.Context, code int) {
	Send(c, http.StatusOK, code, nil)
}

// 自动根据 code 获取 message, 且 data != nil
func SendData(c *gin.Context, code int, data any) {
	Send(c, http.StatusOK, code, data)
}

func SuccessData(c *gin.Context, data any) {
	Send(c, http.StatusOK, OK, data)
}

func Success(c *gin.Context) {
	Send(c, http.StatusOK, OK, nil)
}
