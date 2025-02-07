package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/zhangyiming748/MultiTranslatorUnifier/bootstrap"
	"github.com/zhangyiming748/MultiTranslatorUnifier/model"
	"github.com/zhangyiming748/MultiTranslatorUnifier/storage"
	"github.com/zhangyiming748/MultiTranslatorUnifier/util"

	"github.com/gin-gonic/gin"
)

func testResponse(c *gin.Context) {
	c.JSON(http.StatusGatewayTimeout, gin.H{
		"code": http.StatusGatewayTimeout,
		"msg":  "timeout",
	})
}

func init() {
	//util.SetLog() // 移除这行
	storage.ConnectToMySQL()
	storage.InitRedis()
	storage.GetMysql().Sync2(new(model.TranslateHistory))
}

func main() {
	// if err := storage.GetMysql().Sync2(new(model.TranslateHistory)); err != nil {
	// 	log.Printf("同步MySQL数据表出错%v\n", err)
	// }

	// gin服务
	gin.SetMode(gin.DebugMode)

	// 在 SetLog() 中获取 MultiWriter
	multiWriter := util.SetLog()

	// 将 Gin 的日志输出设置为 MultiWriter
	gin.DefaultWriter = multiWriter

	engine := gin.New()
	// 自定义 Logger 格式
	engine.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[CUSTOM] %s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	bootstrap.InitTranslate(engine)

	h := new(model.TranslateHistory)
	h.Src = "1"
	h.Dst = "2"
	h.InsertOne()
	// 启动http服务
	err := engine.Run(":8192")
	if err != nil {
		log.Fatalln("gin服务启动失败,当前端口有可能被占用")
	}
}
