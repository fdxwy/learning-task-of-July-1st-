package main

import (
	"fmt"
	//导入http的基础封装和访问
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"hello world")
}

func main() {
	//在根路由 / 上注册一个 indexHandler
	http.HandleFunc("/", indexHandler)
	//开启监听
	http.ListenAndServe(":8803", nil)
}
