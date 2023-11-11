package main

import (
	"golang.org/x/sync/errgroup"
	"log"
	"outdoor_rental/dao"
	"outdoor_rental/routes"
	"outdoor_rental/utils"
)

var g errgroup.Group

func main() {
	//初始化viper
	utils.InitViper()
	//初始化数据库
	dao.DB = utils.InitMySQLDB()
	// 初始化 Logger
	utils.InitLogger()
	//初始化redis
	utils.InitRedis()

	// 后台接口服务
	g.Go(func() error {
		return routes.AdminServer().ListenAndServe()
	})
	// 前台接口服务
	g.Go(func() error {
		return routes.FrontServer().ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
