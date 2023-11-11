package utils

import (
	"github.com/spf13/viper"
	"log"
	"outdoor_rental/config"
	"strings"
)

func InitViper() {
	configPath := "config/config.toml"
	// 目前读取固定固定路径的配置文件
	v := viper.New()
	v.SetConfigFile(configPath)
	v.AutomaticEnv()                                   // 允许使用环境变量
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // SERVER_APPMODE => SERVER.APPMODE 创建了一个字符串替换器，它会将所有的点号 .

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		log.Panic("配置文件读取失败: ", err)
	}
	// 加载配置文件内容到结构体对象
	if err := v.Unmarshal(&config.Cfg); err != nil {
		log.Panic("配置文件内容加载失败: ", err)
	}
	log.Println("配置文件内容加载成功")
}
