package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数
	fmt.Println(r.Form) //在服务端打印请求参数
	fmt.Println("URL:", r.URL.Path) //请求URL
	fmt.Println("Scheme:", r.URL.Scheme)

	for k, v := range r.Form {
		fmt.Println(k, ":", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "你好，世界！") //发送响应到客户端
}

func main() {
	http.HandleFunc("/", sayHelloWorld)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe：", err)
	}
}
