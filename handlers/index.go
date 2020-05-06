package handlers

import (
	"hello-go/models"
	"html/template"
	"log"
	"net/http"
)

// 论坛首页路由处理器方法
func Index(w http.ResponseWriter, r *http.Request) {
	files := []string{"views/layout.html", "views/navbar.html", "views/index.html",}

	// 使用 Go 自带的 html/template 作为模板引擎
	templates := template.Must(template.ParseFiles(files...))

	// 从数据库查询群组数据并将该数据传递到模板文件，最后将模板视图渲染出来
	threads, err := models.Threads();

	if err == nil {
		log.Println("打印数据", threads)
		templates.ExecuteTemplate(w, "layout", threads)
	}else{
		log.Fatalln("出现异常", err)
	}
}