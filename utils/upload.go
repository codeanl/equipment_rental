package utils

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"outdoor_rental/utils/r"
)

func Upload(c *gin.Context) (code int, url string) {
	var AccessKey = "dXjfJ47I2-qY5SCiQ2KvWENU8BXsTXKiNpvocA9I"
	var SerectKey = "VnWaBZkO9_AUcsEjr7iZd8-XqYn7nEUlKLMx0fFO"
	var Bucket = "outdoolhuwai"                 //仓库名
	var ImgUrl = "s40hitqos.hn-bkt.clouddn.com" //域名

	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		// 处理错误
		return r.FAIL, "1"
	}
	defer file.Close()

	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SerectKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	key := generateTimestampKey()
	err = formUploader.Put(context.Background(), &ret, upToken, key, file, handler.Size, &putExtra)
	if err != nil {
		return r.FAIL, "2"
	}
	imageURL := "http://" + ImgUrl + "/" + ret.Key
	return r.OK, imageURL
}
func generateTimestampKey() string {
	timestamp := time.Now().Unix()
	key := fmt.Sprintf("%d", timestamp)
	return key
}
