package main

import (
	"MultiTranslatorUnifier/bootstrap"
	"MultiTranslatorUnifier/util"
	"fmt"
	"log"
	"net/http"
	"time"

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
	port := ":8192"
	server := &http.Server{
		Addr:         port,
		Handler:      engine,
		ReadTimeout:  0,                 // 禁用 ReadTimeout
		WriteTimeout: 0,                 // 禁用 WriteTimeout
		IdleTimeout:  120 * time.Second, // 保持连接不断开, 喵！
	}

	// 启动服务器
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("gin服务启动失败,当前端口%s有可能被占用：%v", port, err)
	}
	//err := engine.Run(port)
	//if err != nil {
	//	log.Fatalf("gin服务启动失败,当前端口%s有可能被占用", port)
	//}
}
