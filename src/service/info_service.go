package service

import (
	"fmt"

	"github.com/mouday/cron-runner-shell/src/config"
)

func EchoInfo() {

	fmt.Println("********************************************")
	fmt.Println("* Welcome Use Shell Runner", config.VERSION)
	fmt.Println("* server run at: ", "http://"+config.GetAppRunAddress())
	fmt.Println("* issue: https://github.com/mouday/shell-runner")
	fmt.Println("* X-Token: ", config.GetToken())
	fmt.Println("********************************************")

}
