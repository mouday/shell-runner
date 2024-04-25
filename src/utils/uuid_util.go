package utils

import (
	uuid "github.com/satori/go.uuid"
)

// 获取uuid
func GetUuidV4() string {
	return uuid.NewV4().String()
}
