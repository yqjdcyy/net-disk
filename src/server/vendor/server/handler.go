package server

import (
	"fmt"
	"net/http"
	"net/url"

	"bitbucket.org/ansenwork/ilog"
	"config"
	"utils"
)

// UploadHandler 上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {

	// int
	args := r.URL.Query()
	var path string
	var filename string

	// check
	pathes, ok := args["path"]
	if ok && len(pathes) > 0 {

		if p, err := url.QueryUnescape(pathes[0]); nil != err {
			ilog.Errorf("fail to parse path[%s] to string", pathes[0])
			return
		} else {
			path = p
		}
	}
	filenames, ok := args["filename"]
	if ok && len(filenames) > 0 {

		if f, err := url.QueryUnescape(filenames[0]); nil != err {
			ilog.Errorf("fail to parse filename[%s] to string", filenames[0])
			return
		} else {
			filename = f
		}
	}
	if 0 == len(path) || 0 == len(filename) {
		fmt.Fprintf(w, "args[path|filename] not support")
		return
	}

	// save
	p := utils.Generate(config.Gateway.Dir, path, filename)
	ilog.Debugf("generate.path(%s, %s, %s)= %s", config.Gateway.Dir, path, filename, p)
	if !utils.IsExists(p) {
		if err := utils.Save(p, r.Body); nil != err {
			ilog.Errorf("fail to save file[%s]: %s", p, err.Error())
		}
	} else {
		ilog.Infof("File[%s] is exists", p)
	}
}
