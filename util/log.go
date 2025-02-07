package util

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/zhangyiming748/lumberjack"
)

func SetLog() io.Writer { // 修改返回值
	// 创建一个用于写入文件的Logger实例
	fileLogger := &lumberjack.Logger{
		Filename:   strings.Join([]string{"gin.log"}, string(os.PathSeparator)),
		MaxSize:    1, // MB
		MaxBackups: 1,
		MaxAge:     28, // days
		LocalTime:  true,
	}
	fileLogger.Rotate()
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)

	// 创建 MultiWriter
	multiWriter := io.MultiWriter(fileLogger, consoleLogger.Writer())

	log.SetOutput(multiWriter)
	log.SetFlags(log.Ltime | log.Lshortfile)

	return multiWriter // 返回 MultiWriter
}
