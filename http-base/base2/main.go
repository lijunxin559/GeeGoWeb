package main

import(
    "fmt"
    "log"
    "net/http"
)

type Engine struct {}
//engine这个实例的原有接口实现了ServeHTTP，所以就直接能被调用
func (engin * Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    switch req.URL.Path {
    case "/":
        fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
    case "/hello":
        for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
    default:
        fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
    }
}

func main() {
    engine := new(Engine)
    log.Fatal(http.ListenAndServe(":9999",engine))
}