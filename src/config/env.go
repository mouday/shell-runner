package config

import (
	"os"

	"github.com/mouday/cron-runner-shell/src/utils"
	"github.com/subosito/gotenv"
)

// token 路径
const TOKEN_FILE_PATH = "./token.txt"

func init() {
	gotenv.Load()
}

func GetEnv(key string, defaultValue string) string {

	value := os.Getenv(key)

	if value == "" {
		value = defaultValue
	}

	return value
}

func GetGinMode() string {
	return GetEnv("GIN_MODE", "release")
}

// 启动服务端口
func GetAppRunAddress() string {
	return GetEnv("APP_RUN_ADDRESS", "127.0.0.1:8000")
}

func GetAppAdminUsername() string {
	return GetEnv("APP_ADMIN_USERNAME", "admin")
}

func GetAppAdminPassword() string {
	password := os.Getenv("APP_ADMIN_PASSWORD")

	// if password == "" {
	// 	password = utils.GetRandomString(10)
	// }

	return password

}

func GetScriptDir() string {
	return GetEnv("APP_SCRIPT_DIR", "./scripts")
}

func GetToken() string {
	var token string

	// 尝试从文件中读取
	if utils.FileExists(TOKEN_FILE_PATH) {
		content, err := os.ReadFile(TOKEN_FILE_PATH)
		if err == nil {
			token = string(content)
		}
	}

	// 读取失败则新生成一个
	if token == "" {
		token = utils.GetUuidV4()
		// ref: https://zhuanlan.zhihu.com/p/48529974
		// 0644->即用户具有读写权限，组用户和其它用户具有只读权限
		os.WriteFile(TOKEN_FILE_PATH, []byte(token), 0644)
	}

	return token
}
