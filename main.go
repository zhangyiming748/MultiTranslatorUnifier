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
	util.SetLog()
	storage.ConnectToMySQL()
	storage.InitRedis()
	storage.GetMysql().Sync2(model.TranslateHistory{})
}

func main() {
	// gin服务
	gin.SetMode(gin.DebugMode)
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
	// 启动http服务

	err := engine.Run(":8192")
	if err != nil {
		log.Fatalln("gin服务启动失败,当前端口有可能被占用")
	}
}
