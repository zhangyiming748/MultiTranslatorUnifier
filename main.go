package main

import (
	"fmt"
	"log"
	"time"

	"github.com/zhangyiming748/MultiTranslatorUnifier/bootstrap"
	"github.com/zhangyiming748/MultiTranslatorUnifier/util"

	"github.com/gin-gonic/gin"
)

func main() {

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

	// 启动http服务
	err := engine.Run(":8192")
	if err != nil {
		log.Fatalln("gin服务启动失败,当前端口有可能被占用")
	}
}
