package logic

import (
	"MultiTranslatorUnifier/github"
	"MultiTranslatorUnifier/linuxdo"
	translateshell "MultiTranslatorUnifier/translate-shell"
	"log"
	"sync"
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
	result := <-dst
	log.Printf("result = %s\n", result)
	return result
}
