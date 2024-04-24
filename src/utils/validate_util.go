package utils

import (
	"regexp"
)

// 仅允许数字、字母和下划线
// ref: https://www.cnblogs.com/linck/p/10669487.html
func ValidateScriptName(name string) bool {
	matched, _ := regexp.MatchString(`^[0-9a-zA-Z_]{1,}$`, name)
	return matched
}
