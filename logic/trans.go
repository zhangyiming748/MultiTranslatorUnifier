package logic

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/zhangyiming748/MultiTranslatorUnifier/linuxdo"
	"github.com/zhangyiming748/MultiTranslatorUnifier/storage"
	translateshell "github.com/zhangyiming748/MultiTranslatorUnifier/translate-shell"
)

const (
	TIMEOUT = 30
)

func Trans(src string) (from ,ans string) {
	// src := "hello"
	h:=new(storage.History)
	h.Src=src
	if has,_:=h.FindBySrc();has{
		return "cache",h.Dst
	}
	dst := make(chan string, 1)
	//dst := make(chan map[string]string, 1) // 修改为 map[string]string 的通道
	once := new(sync.Once)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	if runtime.GOOS != "windows"{
		if proxy := os.Getenv("PROXY"); proxy != "" {
			go translateshell.TransByGoogle(src, proxy, once, wg, dst)
		} else {
			go translateshell.TransByBing(src, once, wg, dst)
		}
	}
	os.Setenv("LINUXDO", "DrkwqR4tE3DRyOseVibFah62BJXmcIryt4I9rTtzXTs")
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
	h.Dst=result
	h.InsertOne()
	return "new",result
}
