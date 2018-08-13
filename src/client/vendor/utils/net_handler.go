package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// POST 上传文件
func POST(p, u string) error {

	// log
	fmt.Printf("Post(%v, %v)\n", p, u)

	// init
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// read
	fileWriter, err := bodyWriter.CreateFormFile("file", p)
	if nil != err {
		fmt.Errorf("fail to post file because fail to create file[%v] writer\n", p)
		return err
	}
	file, err := os.Open(p)
	if nil != err {
		fmt.Errorf("fail to open file[%v]\n", p)
		return err
	}
	defer file.Close()
	_, ok := io.Copy(fileWriter, file)
	if nil != ok {
		fmt.Errorf("fail to copy file[%v]\n", p)
		return ok
	}

	// post
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := http.Post(u, contentType, bodyBuf)
	if nil != err {
		fmt.Errorf("fail to post file[%v] to url[%v]\n", p, u)
		return err
	}
	defer resp.Body.Close()

	// post.Response
	respBody, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		fmt.Errorf("fail to get response from post to url[%v]\n", u)
		return err
	}
	fmt.Printf("success to post file[%v] to url[%v]: resp.status= [%v], resp.body= [%v]\n", p, u, resp.Status, string(respBody))
	return nil
}
