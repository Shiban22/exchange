package main

import (
	"exchange/config"
	"exchange/router"
)

func main() {
	config.InitConfig()
	r := router.SetAuthRouter()
	config.InitDB()
	r.Run(config.AppConfig.App.Port) // 监听并在 0.0.0.0:8080 上启动服务
}
