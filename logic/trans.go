package logic

import (
	"fmt"
	"github.com/zhangyiming748/MultiTranslatorUnifier/linuxdo"
	translateshell "github.com/zhangyiming748/MultiTranslatorUnifier/translate-shell"
	"os"
	"sync"
	"time"
)

const (
	TIMEOUT = 30
)

func Trans(src string) string {
	// src := "hello"
	dst := make(chan string, 1)
	//dst := make(chan map[string]string, 1) // 修改为 map[string]string 的通道
	once := new(sync.Once)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	if proxy := os.Getenv("PROXY"); proxy != "" {
		go translateshell.TransByGoogle(src, proxy, once, wg, dst)
	} else {
		go translateshell.TransByBing(src, proxy, once, wg, dst)
	}
	if apikey := os.Getenv("LINUXDO"); apikey != "" {
		go linuxdo.TransByLinuxdoDeepLX(src, apikey, once, wg, dst)
	}
	var result string
	select {
	case result = <-dst:
		//constant.Info(fmt.Sprintf("收到翻译结果:%v\n", dst))
	case <-time.After(TIMEOUT * time.Second): // 设置超时时间为30秒
		fmt.Printf("翻译超时,重试\n此时的src = %v\n", src)
		Trans(src)
	}
	return result
}
