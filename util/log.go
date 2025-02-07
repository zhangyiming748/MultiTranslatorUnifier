package util

import (
	"github.com/zhangyiming748/lumberjack"
	"io"
	"log"
	"os"
	"strings"
)

func SetLog() {
	// 创建一个用于写入文件的Logger实例
	fileLogger := &lumberjack.Logger{
		Filename:   strings.Join([]string{"/log/gin.log"}, string(os.PathSeparator)),
		MaxSize:    1, // MB
		MaxBackups: 1,
		MaxAge:     28, // days
		LocalTime:  true,
	}
	fileLogger.Rotate()
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)
	log.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Writer()))
	log.SetFlags(log.Ltime | log.Lshortfile)
}
