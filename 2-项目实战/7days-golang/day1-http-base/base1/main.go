package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	////http.HandleFunc("/", indexHandler)：这行代码设置了一个路由处理器，
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// 当用户访问网站的根路径（例如 "http://localhost:9999/"）时，会调用 indexHandler 函数。
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
