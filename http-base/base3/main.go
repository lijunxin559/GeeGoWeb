package main

import (
    "fmt"
    "net/http"
    "gee"
)

func main() {
    r := gee.New()//设计使用New()创建一个gee的实例
    //设计使用Get()方法添加路由
    r.GET("/",func(w http.ResponseWriter,req *http.Request){
        fmt.Fprintf(w,"URL.Path= %q\n",req.URL.Path)
    })
    r.GET("/hello",func(w http.ResponseWriter,req *http.Request){
        for k,v := range req.Header {
            fmt.Fprintf(w,"Header[%q] = %q\n", k, v)
        }
    })
    //设计使用Run方法启动Web服务
    r.Run(":9000")
}