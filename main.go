package main

import (
	"chat/model"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

var logFile *os.File
var file string

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.SetPrefix("bmc info: ")
	os.Mkdir("log", os.ModePerm)
	file := "./log/" + time.Now().Format("20180102") + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Print(err)
	}
	log.SetOutput(logFile)
}


func main() {

	engine := gin.Default()
	// 注册模板
	engine.LoadHTMLGlob("static/*")
	// 注册静态文件
	engine.Static("./static", "static") // (relativePath, root string): 路径, 文件夹名称
	engine.GET("/chat", model.AjaxTest)	// 获取请求页面
	engine.GET("/", model.AjaxTest)
	engine.POST("/post_ajax", model.PostAjax)

	engine.Run(":9000")


}
