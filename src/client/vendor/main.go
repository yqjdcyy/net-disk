package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"config"
	"logger"
	"utils"
)

var ()

func main() {

	// init
	flag.Parse()
	config.Init()
	logger.SetLogPath(config.ClientConf.LogPath)
	defer logger.Close()

	// jog
	go job()

	// hold
	for {
		select {
		case <-config.Quit:
			logger.Info("client.finish")
			return
		}
	}
}

func job() {

	files := make([]string, 0)
	args := make([]string, 0)
	errs := make([]int, 0)

	for _ = range time.NewTicker(time.Second * time.Duration(config.ClientConf.Duration)).C {

		if 0 == len(config.ClientConf.Path) {
			config.Quit <- true
		}

		// filter
		for _, p := range config.ClientConf.Path {

			_ = filepath.Walk(p, func(q string, file os.FileInfo, err error) error {

				if nil == file {
					return err
				}

				if check(p, q) {
					files = append(files, q)
					args = append(args, q[len(p):])
				}

				return nil
			})
		}

		// handle
		fmt.Println(len(files))
		for i, v := range files {

			p := args[i]
			fs := "\\"
			if runtime.GOOS != "windows" {
				fs = "/"
			}
			idx := strings.LastIndex(p, fs)
			if -1 == idx {
				idx = len(p) - 1
			}

			// upload
			err := utils.POST(v, fmt.Sprintf(config.ClientConf.URL, url.QueryEscape(p[:idx]), url.QueryEscape(p[idx+1:])))
			if nil != err {
				logger.Errorf("fail to upload file[%v]: %v", p, err.Error())
				errs = append(errs, i)
			}
		}

		// remove
		if config.ClientConf.Delete {
			for i, v := range files {

				if !contain(errs, i) {
					utils.Remove(v)
				}
			}
		}

		// quit
		config.Quit <- true
	}
}

func check(p, q string) bool {

	// TODO yqj check.hour
	if utils.ContainType(config.ClientConf.Type, utils.GetExtension(q)) && !utils.IsFolder(q) {
		return true
	}
	return false
}

func contain(s []int, o int) bool {

	for _, v := range s {
		if v == o {
			return true
		}
	}
	return false
}
