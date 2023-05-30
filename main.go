package main

import (
	"chat/model"
	"github.com/gin-gonic/gin"
	//"io"
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
	//db,_ := libs.InitDB()
	//lb_uuid,bool_res,_ := model.EsLBCheck(db)
	//if bool_res {
	//
	//}

	//var ss string = model.ChatGpt("arista交换机命令")
	//log.Print("1111"+ss)

	//r := gin.Default() //1.指定默认路由
	//r.Use(cors.New(cors.Config{
	//	AllowOrigins:  []string{"http://127.0.0.1:8080", "http://127.0.0.1:8000"},
	//	AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	//	//AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
	//}))
	//r.GET("/api/lbverfyres", model.LbVerfyres)
	//r.GET("/api/zoneverfy", model.SpecialZoneLbVerfy)
	////r.GET("/api/ai", model.ChatGpt) {}
	//r.Run(":8081") //3.监听端口

	engine := gin.Default()
	// 注册模板
	engine.LoadHTMLGlob("static/*")
	// 注册静态文件
	engine.Static("./static", "static") // (relativePath, root string): 路径, 文件夹名称
	engine.GET("/chat", model.AjaxTest)	// 获取请求页面
	engine.POST("/post_ajax", model.PostAjax)

	engine.Run(":9001")


}
