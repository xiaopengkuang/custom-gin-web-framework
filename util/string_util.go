package util

import "strings"

const (
	EmptyString = ""
)

// 判断字符串是否为空
func IsEmpty(str string) bool {
	return strings.TrimSpace(str) == EmptyString
}

// 是否全有字符
func DoAllArgsHasText(args ...string) bool {
	if args == nil || len(args) == 0 {
		return false
	}

	for _, arg := range args {
		if IsEmpty(arg) {
			return false
		}
	}

	return true
}
