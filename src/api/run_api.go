package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mouday/cron-runner-shell/src/config"
	"github.com/mouday/cron-runner-shell/src/utils"
	"github.com/mouday/cron-runner-shell/src/vo"
)


/*
 * 登录
 */
func RunScript(ctx *gin.Context) {
	name := ctx.Query("name")
	token := ctx.Request.Header.Get("X-Token")

	if name == "" || token == "" || token != config.GetToken() {
	    vo.Error(ctx, -1, "token无效")
	} else {
	// 异步运行
	go utils.RunShellScript("scripts", name)

	vo.Success(ctx, nil)
	}

}
