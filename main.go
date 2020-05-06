package main

import (
	. "hello-go/routes"
	"log"
	"net/http"
)

func main()  {
	startWebServer("8080")
}

// 通过指定端口启动 Web 服务器
func startWebServer(port string)  {
	r := NewRouter() // 通过 router.go 中定义的路由器来分发请求

	// 处理静态资源文件
	// 将 /static/ 前缀的 URL 请求去除 static 前缀，
	// 然后在文件服务器查找指定文件路径是否存在（public 目录下的相对地址）。
	assets := http.FileServer(http.Dir("public"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r) // 应用路由器到 HTTP 服务器

	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":" + port, nil) // 启动协程监听请求

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}