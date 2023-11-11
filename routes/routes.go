package routes

import (
	"log"
	"net/http"
	"outdoor_rental/config"
	"time"
)

// 后台服务
func AdminServer() *http.Server {
	backPort := config.Cfg.Server.BackPort
	log.Printf("后台服务启动于 %s 端口", backPort)
	return &http.Server{
		Addr:         backPort,
		Handler:      AdminRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

// 前台服务
func FrontServer() *http.Server {
	FrontPort := config.Cfg.Server.FrontPort
	log.Printf("前台服务启动于 %s 端口", FrontPort)
	return &http.Server{
		Addr:         FrontPort,
		Handler:      FrontRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
