package handlers

import (
	"hello-go/models"
	"log"
	"net/http"
)

// 论坛首页路由处理器方法
func Index(writer http.ResponseWriter, request *http.Request) {
	// 从数据库查询群组数据并将该数据传递到模板文件，最后将模板视图渲染出来
	threads, err := models.Threads();

	if err == nil {
		//log.Println("打印数据", threads)
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, threads, "layout", "navbar", "index")
		} else {
			generateHTML(writer, threads, "layout", "auth.navbar", "index")
		}
	}else{
		log.Fatalln("出现异常", err)
	}
}

func Err(writer http.ResponseWriter, request *http.Request)  {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "auth.navbar", "error")
	}
}