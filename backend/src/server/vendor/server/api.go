package server

import (
	"fmt"
	"net/url"
	"net/http"
	"path/filepath"
	"encoding/json"
	"strconv"

	"config"
	// "bitbucket.org/ansenwork/ilog"
)

// FileHandler 界面相关
func FileHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method{
		case http.MethodGet: 
			pageFiles(w, r)
		case http.MethodDelete:
			fmt.Fprintf(w, "todo")
	default:
		fmt.Fprintf(w, "Method["+ r.Method+ "] is not support");
	}

}

// pageFiles 分页查询
func pageFiles(w http.ResponseWriter, r *http.Request){

	// init
	path, page, size:= listPageArgs(r.URL.Query())

	// check
	if 0== len(path){
		http.Error(w, "without arg[path]", http.StatusInternalServerError)
		return
	}
	
	// query
	// todo: cache
	var cache []string
	if 0== len(cache){
		
		cache = listFiles(path)
		// todo: cache
	}

	// filter
	list:= subArr(cache, page, size)

	// response
	bytes, err := json.Marshal(&APIResponse{page, size , len(cache), config.Gateway.Suffix, list})
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func listPageArgs(args url.Values) (path string, page int, size int){

	pathes, ok := args["path"]
	if ok && len(pathes) > 0 {
		path= pathes[0]
	}

	page= 1
	pages, ok := args["page"]
	if ok && len(pages)> 0{
		page, _= strconv.Atoi(pages[0])
	}

	size= 10
	sizes, ok := args["size"]
	if ok && len(sizes)> 0{
		size, _= strconv.Atoi(sizes[0])
	}

	return
}

func listFiles(path string) []string {

	list, err := filepath.Glob(filepath.Join(path,"*"))
	if err != nil {
		fmt.Println(err)
		list= []string{}
	}

	return list
}

func subArr(list []string, page int, size int) []string {

	count:= len(list)
	if 0== count{
		return list
	}

	from:= (page-1)* size
	if from < 0{
		from= 0
	}
	if from> count{
		return []string{}
	}
	to := page* size
	if to> count{
		to = count
	}
	return list[from:to]
}