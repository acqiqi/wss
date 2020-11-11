package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime"
	"wss/lib/gredis"
	"wss/lib/logging"
	"wss/lib/setting"
	"wss/lib/utils"
	"wss/models"
	"wss/router"
	_ "wss/task"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) //设置全局log 打印带行数
	log.Println("Init Project")
	setting.Setup() //处理配置文件
	models.Setup()  //模型
	logging.Setup() //日志
	gredis.Setup()  // redis
	//utils.InitModel()
	//utils.PDFT()
	log.Println(utils.FormatCamelString("asd_asd", false))

}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := router.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	log.Println(runtime.Version())
	server.ListenAndServe()
}
