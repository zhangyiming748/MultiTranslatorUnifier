package linuxdo

import (
	"encoding/json"
	"fmt"
	"github.com/zhangyiming748/MultiTranslatorUnifier/util"
	"log"
	"strings"
	"sync"
)

const PREFIX = "https://api.deeplx.org"
const SUFFIX = "translate"

type DeepLXTranslationResult struct {
	Code         int      `json:"code"`
	ID           int64    `json:"id"`
	Message      string   `json:"message,omitempty"`
	Data         string   `json:"data"`         // The primary translated text
	Alternatives []string `json:"alternatives"` // Other possible translations
	SourceLang   string   `json:"source_lang"`
	TargetLang   string   `json:"target_lang"`
	Method       string   `json:"method"`
}

func TransByLinuxdoDeepLX(src, apikey string, once *sync.Once, wg *sync.WaitGroup, dst chan string) {
	//apikey := os.Getenv("LINUXDO")
	result, err := Req(src, apikey)
	result = strings.Replace(result, "\\r\\n", "", 1)
	result = strings.Replace(result, "\n", "", 1)
	result = strings.Replace(result, "\r\n", "", 1)
	if result == "" {
		return
	}
	log.Printf("linuxdo 版本 deeplx 返回:%+v\n", result)
	if err != nil {
		log.Printf("linuxdo 版本 deeplx 查询执行出错\t错误原文:%v\n", err.Error())
		return
	} else {
		once.Do(func() {
			fmt.Println("Bing返回翻译结果")
			dst <- result
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
	var d DeepLXTranslationResult
	if e := json.Unmarshal(b, &d); e != nil {
		return "", e
	}
	return d.Data, err
}
