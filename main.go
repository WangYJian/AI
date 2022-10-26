package main

import (
	"AI/conf"
	"AI/Key"
	"AI/router"
	"github.com/gin-gonic/gin"
)
func SettingUpEnvironment() {
	//读入配置文件
	c := conf.ReadSettingsFromFile("Config.json")

	//配置信息
	Key.InitKey(c.Settings)

}

func main() {
	//设置初始环境
	SettingUpEnvironment()
	
	//创建路由对象
	f := gin.Default()

	//使用路由功能
	router.UseMyRouter(f)

	//启动
	f.Run(":8080");

}