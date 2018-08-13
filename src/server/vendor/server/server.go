package server

import (
	"config"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Start 服务启动
func Start() {

	r := mux.NewRouter()
	r.HandleFunc("/upload", UploadHandler).Methods("POST")

	http.ListenAndServe(fmt.Sprintf(":%d", config.Gateway.Port), r)
}
