package utils

import (
	"errors"
	"log"
	"strings"
)

// Parse 参数格式化处理
func Parse(path string) (*map[string]string, error) {

	// check
	if len(strings.Trim(path, "")) == 0 {
		return nil, errors.New("path is empty")
	}
	idx := strings.Index(path, "=")
	if idx < 0 {
		return nil, errors.New("abnormal format path")
	}

	// init
	argsMap := map[string]string{}
	idx = strings.Index(path, "?")
	if idx >= 0 {
		path = path[idx+1:]
	}

	// split+ fill
	args := strings.Split(path, "&")
	for _, arg := range args {

		a := strings.Split(arg, "=")
		if 2 == len(a) {
			argsMap[a[0]] = a[1]
		} else {
			log.Printf("fail to parse [%v].[%v]\n", path, arg)
		}
	}

	// return
	return &argsMap, nil
}
