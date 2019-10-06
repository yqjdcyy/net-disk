package utils

import (
	"strings"
)

// ContainType 判断是否包含指定类型
func ContainType(ts []string, t string) bool {

	if 0 == len(ts) || 0 == len(t) {
		return false
	}
	t = strings.ToLower(t)

	for _, v := range ts {
		if v == t {
			return true
		}
	}

	return false
}
