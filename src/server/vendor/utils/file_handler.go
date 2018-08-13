package utils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"bitbucket.org/ansenwork/ilog"
)

// OpenOrCreate 打开&创建
func OpenOrCreate(path string) (file *os.File, err error) {

	file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if nil != err {
		ilog.Errorf("OpenOrCreate fail: %v", err.Error())
		return nil, err
	}

	return file, nil
}

// IsExists 判断文件是存在
func IsExists(path string) bool {

	if _, err := os.Stat(path); nil != err && os.IsNotExist(err) {
		return false
	}

	return true
}

// Generate 生成路径
func Generate(d, p, f string) string {

	// remove file.seperator
	stIdx := 0
	if isFileSeperator(p[stIdx]) {
		stIdx = stIdx + 1
	}
	endIdx := len(p)
	if isFileSeperator(p[endIdx-1]) {
		endIdx = endIdx - 1
	}
	p = p[stIdx:endIdx]

	stIdx = 0
	if isFileSeperator(p[stIdx]) {
		stIdx = stIdx + 1
	}
	f = f[stIdx:]

	// system.compatible
	fs := "/"
	ts := "\\"
	if runtime.GOOS != "windows" {
		fs = "\\"
		ts = "/"
	}
	p = strings.Replace(p, fs, ts, -1)

	// format
	return fmt.Sprintf("%s%s%s%s%s", d, ts, p, ts, f)
}

func isFileSeperator(s byte) bool {

	if '/' == s || '\\' == s {
		return true
	}
	return false
}

// GetExtension 获取文件后缀名
func GetExtension(path string) string {

	if 0 == len(path) {
		return ""
	}

	idx := strings.LastIndex(path, ".")
	if -1 == idx {
		return ""
	}

	return path[idx+1:]
}

// IsFolder 判断指定路径是否为文件
func IsFolder(path string) bool {

	fi, err := os.Stat(path)
	if nil != err && os.IsNotExist(err) {
		return false
	}

	return fi.IsDir()
}

// Contain 检测文件是否包含指定字符串
func Contain(path, key string) bool {

	if 0 == len(path) || 0 == len(key) {
		return false
	}

	if !IsExists(path) {
		return false
	}

	f, err := os.Open(path)
	if nil != err {
		ilog.Errorf("fail to check Contain(%v): %v", path, err.Error())
		return false
	}

	bs, err := ioutil.ReadAll(f)
	if nil != err {
		ilog.Errorf("fail to check Contain(%v): %v", path, err.Error())
		return false
	}
	if strings.Index(string(bs), key) > 0 {
		return true
	}

	return false
}

// Save 将文件保存至指定目录文件
func Save(filePath string, reader io.ReadCloser) (err error) {

	defer reader.Close()

	// dir.create
	path := filepath.Dir(filePath)
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Printf("fail to mkdir path[%v]: %v", path, err.Error())
		return
	}

	// file.create
	f, err := os.Create(filePath)
	if nil != err {
		fmt.Printf("fail to create file[%v]\n", filePath)
		return
	}
	defer f.Close()

	// file.fill
	bufReader := bufio.NewReader(reader)
	_, err = bufReader.WriteTo(f)
	ilog.Infof("save file[%s]", filePath)
	return
}
