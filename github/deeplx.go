package github

import (
	"fmt"
	"log"
	"sync"

	"github.com/OwO-Network/DeepLX/translate"
)

func TransByGithubDeepLX(src, proxy string, once *sync.Once, wg *sync.WaitGroup, dst chan string) {
	ret, err := translate.TranslateByDeepLX("auto", "zh", src, "", proxy, "")
	log.Printf("GitHub 版本 deeplx 返回:%+v\n", ret)
	result := ret.Data
	if err != nil {
		log.Printf("GitHub 版本 deeplx 查询执行出错\t错误原文:%v\n", err.Error())
		return
	} else {
		once.Do(func() {
			fmt.Println("Bing返回翻译结果")
			dst <- result
			wg.Done()
		})
	}
}
