package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mouday/cron-runner-shell/src/api"
)

/* 注册路由 */
func RegistRouter(app *gin.Engine) {

	app.POST("/api/runScript", api.RunScript)

}
