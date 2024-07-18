package main

import (
	"goBlog/common"
	"goBlog/router"
	"log"
	"net/http"
)

func init() {
	//模板加载
	common.LoadTemplate()
}
func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Panicln(err)
	}
}
//gengxin
