package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mouday/cron-runner-shell/src/config"
)

const VERSION = "v1.0.0"

func main() {
	// app
	env := config.GetEnv()
	if env == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.New()
	app.SetTrustedProxies(nil)

	appRunAddress := config.GetAppRunAddress()

	fmt.Println("********************************************")
	fmt.Println("* Welcome Use Shell Runner", VERSION)
	fmt.Println("* server run at: ", "http://"+appRunAddress)
	fmt.Println("* issue: https://github.com/mouday/cron-admin")
	fmt.Println("********************************************")

	// 监听并在 http://127.0.0.1:8082 上启动服务
	err := app.Run(appRunAddress)

	if err != nil {
		fmt.Println(err)
	}
}
