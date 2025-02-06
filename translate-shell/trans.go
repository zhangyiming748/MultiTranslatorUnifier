package translateshell

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"
)

func TransByGoogle(src, proxy string, once *sync.Once, wg *sync.WaitGroup, dst chan map[string]string) {
	cmd := exec.Command("trans", "-brief", "-engine", "google", "-proxy", proxy, ":zh-CN", src)
	output, err := cmd.CombinedOutput()
	result := string(output)
	m := map[string]string{
		"Google": result,
	}
	if err != nil || strings.Contains(string(output), "u001b") || strings.Contains(string(output), "Didyoumean") || strings.Contains(string(output), "Connectiontimedout") {
		log.Printf("google查询命令执行出错\t命令原文:%v\t错误原文:%v\n", cmd.String(), err.Error())
		return
	} else {
		once.Do(func() {
			fmt.Println("Google返回翻译结果")
			dst <- m
			wg.Done()
		})
	}
}

func TransByBing(src, proxy string, once *sync.Once, wg *sync.WaitGroup, dst chan map[string]string) {
	cmd := exec.Command("trans", "-brief", "-engine", "bing", "-proxy", proxy, ":zh-CN", src)
	log.Printf("查询命令:%s\n", cmd.String())
	output, err := cmd.CombinedOutput()
	result := string(output)
	m := map[string]string{
		"Bing": result,
	}
	if err != nil || strings.Contains(string(output), "u001b") || strings.Contains(string(output), "Didyoumean") || strings.Contains(string(output), "Connectiontimedout") {
		log.Printf("bing查询命令执行出错\t命令原文:%v\t错误原文:%v\n", cmd.String(), err.Error())
		return
	} else {
		once.Do(func() {
			fmt.Println("Bing返回翻译结果")
			dst <- m
			wg.Done()
		})
	}
}
