package linuxdo

import (
	"encoding/json"
	"fmt"
	"github.com/zhangyiming748/MultiTranslatorUnifier/util"
	"log"
	"strings"
	"sync"

	public "github.com/OwO-Network/DeepLX/translate"
)

const PREFIX = "https://api.deeplx.org"
const SUFFIX = "translate"

func TransByLinuxdoDeepLX(src, apikey string, once *sync.Once, wg *sync.WaitGroup, dst chan map[string]string) {
	//apikey := os.Getenv("LINUXDO")
	result, err := Req(src, apikey)
	result = strings.Replace(result, "\\r\\n", "", 1)
	result = strings.Replace(result, "\n", "", 1)
	result = strings.Replace(result, "\r\n", "", 1)
	if result == "" {
		return
	}
	log.Printf("linuxdo 版本 deeplx 返回:%+v\n", result)
	m := map[string]string{
		"LinuxDo": result,
	}
	if err != nil {
		log.Printf("linuxdo 版本 deeplx 查询执行出错\t错误原文:%v\n", err.Error())
		return
	} else {
		once.Do(func() {
			fmt.Println("Bing返回翻译结果")
			dst <- m
			wg.Done()
		})
	}
}
func Req(src, apikey string) (string, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	params := map[string]string{
		"text":        src,
		"source_lang": "auto",
		"target_lang": "zh",
	}
	host := strings.Join([]string{PREFIX, apikey, SUFFIX}, "/")

	b, err := util.HttpPostJson(headers, params, host)
	if err != nil {
		return "", err
	}
	log.Printf("%v\n", string(b))
	var d public.DeepLXTranslationResult
	if err := json.Unmarshal(b, &d); err != nil {
		return "", err
	}
	return d.Data, err
}
