package logic

import (
	"MultiTranslatorUnifier/github"
	"MultiTranslatorUnifier/linuxdo"
	translateshell "MultiTranslatorUnifier/translate-shell"
	"fmt"
	"log"
	"sync"
	"time"
)

const (
	TIMEOUT = 30
)

func Trans(src, proxy, apikey string) map[string]string {
	// src := "hello"
	// proxy := "http://192.168.1.3:8889"
	// dst := make(chan string, 1)
	dst := make(chan map[string]string, 1) // 修改为 map[string]string 的通道
	once := new(sync.Once)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	if proxy != "" {
		go translateshell.TransByBing(src, proxy, once, wg, dst)
		go translateshell.TransByGoogle(src, proxy, once, wg, dst)
	}
	if apikey != "" {
		go linuxdo.TransByLinuxdoDeepLX(src, apikey, once, wg, dst)
	}
	go github.TransByGithubDeepLX(src, proxy, once, wg, dst)
	var result map[string]string
	select {
	case result = <-dst:
		//constant.Info(fmt.Sprintf("收到翻译结果:%v\n", dst))
	case <-time.After(TIMEOUT * time.Second): // 设置超时时间为30秒
		fmt.Printf("翻译超时,重试\n此时的src = %v\n", src)
		Trans(src, proxy, apikey)
	}
	log.Printf("result = %s\n", result)
	return result
}
