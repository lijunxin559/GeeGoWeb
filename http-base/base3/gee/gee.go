package gee

import(
    "fmt"
    "net/http"
)

type HandlerFunc func(w http.ResponseWriter,req *http.Request)

type Engine struct{
    //路由映射表
    router map[string]HandlerFunc
}

func New() *Engine{
    return &Engine{router:make(map[string]HandlerFunc)}
}

func (engine * Engine) addRoute(method string,pattern string,handler HandlerFunc){
    key := method + "-" + pattern
    engine.router[key] = handler
}


func (engine * Engine) GET(pattern string,handler HandlerFunc){
    engine.addRoute("GET",pattern,handler)
}

func (engine * Engine) POST(pattern string,handler HandlerFunc){
    engine.addRoute("POST",pattern,handler)
}

func (engine * Engine) ServeHTTP(w http.ResponseWriter,req *http.Request){
    key := req.Method + "-"  + req.URL.Path
    if handler,ok := engine.router[key];ok{
        handler(w,req)
    }else {
        fmt.Fprintf(w,"404 NOT FOUND: %s\n",req.URL)
    }
}

func (engine * Engine) Run(addr string) (err error){
    return http.ListenAndServe(addr,engine)
}

