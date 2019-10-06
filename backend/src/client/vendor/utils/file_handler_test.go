package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestOpenOrCreate(t *testing.T) {

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

func TestGetExtension(t *testing.T) {

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

func TestGetMMin(t *testing.T) {

	p := "‪D:\\download\\test\\2018.jpg"
	d := 100
	mmin := GetMMin(p)

	t.Logf("\tGetMMIN(%s)= %v\n", p, mmin)

	if mmin > d {
		t.Errorf("File[%v] last[%v] for more than %v min", p, mmin, d)
	}
}

func TestRemove(t *testing.T) {

	p := "‪D:\\download\\test\\2018.jpg"

	err := Remove(p)
	if nil != err {
		// t.Errorf("fail to remove File[%s]: %s", p, err.Error())
		return
	}

	t.Logf("File[%v] removed", p)
}
