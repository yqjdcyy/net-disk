package main

import (
	"fmt"
	"path/filepath"
	"encoding/json"

	"os"
)


func main() {
	testRemoveFiles()
}

func testRemoveFiles(){
	list:= listFiles("D:\\data\\tmp\\img")

	for _, v := range list {
		if deleteFile(v){
			fmt.Println("delete "+ v)
		}else{
			fmt.Println("fail to delete "+ v)
		}
	}
}

func deleteFile(path string) bool{

	bFile:=isFile(path)
	fmt.Println(path )
	fmt.Println( bFile)
	if !bFile{
		// fmt.Println(33333)
		return false
	}

	err:= os.Remove(path)
	if nil!= err{
		fmt.Println("fail to remove file["+ path+"]: "+ err.Error())
		return false
	}
	
	fmt.Println("remove file: "+ path)
	return true
}

func isFile(path string) bool{

	if 0== len(path){
		fmt.Println("without path")
		return false
	}

	info,err := os.Stat(path)
	fmt.Println(err)
	fmt.Println(info.IsDir())
	if nil!= err {
		return false
	}

	return !info.IsDir()
}

func testListFiles(){

	list:= listFiles("E:/image/wallpaper")
	size:= 10
	count:= len(list)/ size
	listAll(subArr(list, count+1, size))
	listAll(subArr(list, count+2, size))

	response(list)
}

func listFiles(path string) []string {

	list, err := filepath.Glob(filepath.Join(path,"*"))
	if err != nil {
		list= []string{"1"}
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

func listAll(list []string){
	
	fmt.Println("--------------------")
	for _, v := range list {
		fmt.Println(v)
	}
}

func response(list []string){
	
	res:= &APIResponse{ 1, 10 , len(list), "_30x30", list}
	bytes, err := json.Marshal(res)
	if err != nil {
	  fmt.Println(err.Error())
	  return
	}
	fmt.Println(string(bytes))
}

//APIResponse response
type APIResponse struct {
	Page int	`json:"page"`
	 Size int	`json:"size"`
	 Count int	`json:"count"`
	Suffix string	`json:"suffix"`
	List []string	`json:"list"`
}