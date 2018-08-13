package utils

import (
	"encoding/json"
)

// ToString 对象转文本
func ToString(v interface{}) string {

	bs, err := json.Marshal(v)
	if nil != err {
		return ""
	}

	return string(bs)
}
