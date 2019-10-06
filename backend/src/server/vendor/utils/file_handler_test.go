package utils

import (
	"fmt"
	"testing"
	"time"
)

func Test_OpenOrCreate(t *testing.T) {

	s := time.Now()
	file, err := OpenOrCreate("/data/logs/yao.log")
	if nil != err {
		fmt.Print(err.Error())
	} else {
		fmt.Print(file.Name())
	}
	d := time.Since(s)
	fmt.Print(d)
}

func Test_GetExtension(t *testing.T) {

	getExtensionSample(t, "", "")
	getExtensionSample(t, "D:/work/file_handler_test", "")
	getExtensionSample(t, "D:/work/file_handler_test.go", "go")
}

func getExtensionSample(t *testing.T, path, extension string) {

	if extension == GetExtension(path) {
		t.Logf("GetExtension(%v) success", path)
	} else {
		t.Errorf("GetExtension(%v) fail", path)
	}
}

func TestIsFolder(t *testing.T) {

	path := "D:\\work"

	if IsFolder(path) {
		t.Logf("IsFolder(%v) correct", path)
	} else {
		t.Errorf("IsFolder(%v) fail", path)
	}
}

func TestContain(t *testing.T) {

	path := "D:\\work\\git\\yk\\go\\av-server\\src\\av-server\\vendor\\test\\correct-index.m3u8"
	key := "#EXT-X-ENDLIST"

	if Contain(path, key) {
		t.Logf("Contain(%v, %v) correct", path, key)
	} else {
		t.Errorf("Contain(%v, %v) fail", path, key)
	}

	path = "D:\\work\\git\\yk\\go\\av-server\\src\\av-server\\vendor\\test\\error-index.m3u8"

	if !Contain(path, key) {
		t.Logf("Contain(%v, %v) correct", path, key)
	} else {
		t.Errorf("Contain(%v, %v) fail", path, key)
	}
}

func TestGenerate(t *testing.T) {

	p = "/path/"
	f := "/file"
	t.Logf(Generate(p, f))
}
