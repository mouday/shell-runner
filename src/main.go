package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mouday/cron-runner-shell/src/config"
	"github.com/mouday/cron-runner-shell/src/router"
	"github.com/mouday/cron-runner-shell/src/utils"
)

func main() {
	// app
	env := config.GetGinMode()
	if env == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.New()
	app.SetTrustedProxies(nil)

	// 注册路由
	router.RegistRouter(app)

	// 启动消费者，确保单线程执行
	go utils.Consumer()

	// 打印版本信息
	utils.EchoInfo()

	// 监听并在 http://127.0.0.1:8082 上启动服务
	err := app.Run(config.GetAppRunAddress())

	if err != nil {
		fmt.Println(err)
	}
}
