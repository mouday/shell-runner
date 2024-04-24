package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mouday/cron-runner-shell/src/config"
	"github.com/mouday/cron-runner-shell/src/utils"
	"github.com/mouday/cron-runner-shell/src/vo"
)

/*
 * 运行脚本
 */
func RunScript(ctx *gin.Context) {
	name := ctx.Query("name")
	token := ctx.GetHeader("X-Token")

	// 校验token
	if token == "" || token != config.GetToken() {
		vo.Error(ctx, -1, "token无效")
		return
	}

	// 校验脚本名称
	if name == "" || !utils.ValidateScriptName(name) {
		vo.Error(ctx, -1, "脚本名称不合法")
		return
	}

	// 校验脚本是否存在
	if !utils.CheckScriptExists(name) {
		vo.Error(ctx, -1, "脚本不存在")
		return
	}

	// 异步运行
	go utils.RunShellScript(name)

	vo.Success(ctx, nil)

}
