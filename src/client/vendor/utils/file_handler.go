package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// OpenOrCreate 打开&创建
func OpenOrCreate(path string) (file *os.File, err error) {

	file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if nil != err {
		fmt.Errorf("OpenOrCreate fail: %v", err.Error())
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

	// format
	return fmt.Sprintf("%s/%s/%s", d, p, f)
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
		fmt.Errorf("fail to check Contain(%v): %v", path, err.Error())
		return false
	}

	bs, err := ioutil.ReadAll(f)
	if nil != err {
		fmt.Errorf("fail to check Contain(%v): %v", path, err.Error())
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
	err = os.MkdirAll(path, os.ModeDir)
	if err != nil {
		fmt.Errorf("fail to mkdir path[%v]: %v", path, err.Error())
		return
	}

	// file.create
	f, err := os.Create(filePath)
	if nil != err {
		fmt.Errorf("fail to create file[%v]\n", filePath)
		return
	}
	defer f.Close()

	// file.fill
	bufReader := bufio.NewReader(reader)
	_, err = bufReader.WriteTo(f)
	fmt.Printf("save file[%s]", filePath)
	return
}

// GetMMin 获取文件的最近修改时间
func GetMMin(p string) int {

	if 0 == len(p) {
		return 0
	}

	info, err := os.Stat(p)
	if nil != err {
		fmt.Errorf("\nfail to read file[%s]: %s\n", p, err.Error())
		return 0
	}
	if nil == info {
		fmt.Errorf("\nfail to get file[%s]\n", p)
		return 0
	}
	return int((time.Now().Unix() - info.ModTime().Unix()) / 60)
}

// Remove 删除文件或目录
func Remove(p string) error {

	if 0 == len(p) {
		return errors.New("File.path is not pointed")
	}

	err := os.Remove(p)
	if nil != err {
		return err
	}

	fmt.Printf("remove File[%v]\n", p)
	return nil
}
