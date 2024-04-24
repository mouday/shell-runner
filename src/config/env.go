package config

import (
	"os"

	"github.com/subosito/gotenv"
)

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

func GetToken() string {
	return os.Getenv("APP_TOKEN")
}

func GetScriptDir() string {
	return GetEnv("APP_SCRIPT_DIR", "./scripts")
}
