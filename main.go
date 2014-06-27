package main

import (
	"contacts/controllers"
	"contacts/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

/*
var staticHandler http.Handler

func init() {
	dir := path.Dir(os.Args[0])
//	fmt.Println("dir:", http.Dir(dir))
	staticHandler = http.FileServer(http.Dir(dir))
}

// 静态文件处理
func StaticServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("path:" + req.URL.Path)
	if req.URL.Path != "/down/" {
		staticHandler.ServeHTTP(w, req)
		return
	}

	io.WriteString(w, "hello, world!\n")
}
*/
func main() {
	beego.Get("/download", func(ctx *context.Context) {
		buffer := models.WriteToVCF()
		ctx.Output.Body(buffer.Bytes())
		//	StaticServer(ctx.ResponseWriter, ctx.Request)
	})
	//	beego.Handler("/download", http.FileServer(http.Dir(wd)))

	beego.Router("/", &controllers.MainController{})

	beego.Run()
}
