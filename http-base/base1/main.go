package main

import (
	"fmt"
	"net/http"
	"log"
)

func main()  {
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/hello",helloHandler)
	//设置监听端口，nil代表处理所有的HTTP请求的实例，使用标准库中的实例处理
	log.Fatal(http.ListenAndServe(":9999",nil))
}

//设置两个路由/和/hello，分别绑定下面两个函数
//根据不同的http请求调用不同的处理函数，返回响应

//handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
