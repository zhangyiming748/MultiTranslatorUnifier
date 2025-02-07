package controller

import (
	"fmt"
	"log"

	"github.com/zhangyiming748/MultiTranslatorUnifier/logic"
	"github.com/zhangyiming748/MultiTranslatorUnifier/model"
	"github.com/zhangyiming748/MultiTranslatorUnifier/storage"

	"github.com/gin-gonic/gin"
)

type TranslateController struct{}

/*
curl --location --request GET 'http://127.0.0.1:8192/api/v1/translate?user=zen'
*/
func (t TranslateController) GetTranslate(ctx *gin.Context) {
	user := ctx.Query("user")
	ctx.String(200, fmt.Sprintf("%s还活着!", user))
}

// 结构体必须大写 否则找不到
type RequestBody struct {
	Src        string `json:"src"`
	Proxy      string `json:"proxy"`
	LinuxDoKey string `json:"linuxdokey"`
}

type ResponseBody struct {
	From string `json:"from"`
	Dst  string `json:"dst"`
}

/*
curl --location --request POST 'http://127.0.0.1:8192/api/v1/translate' \
--header 'Content-Type: application/json' \

	--data-raw '{
	    "src":"",
	    "proxy":"",
	    "linuxdokey":""
	}'
*/
func (t TranslateController) PostTranslate(ctx *gin.Context) {
	log.Println("接收到post请求")
	var requestBody RequestBody
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		log.Printf("成功解析post的json:%+v\n", requestBody)
	}
	log.Printf("src:%s\tproxy:%s\tlinuxdo:%s\n", requestBody.Src, requestBody.Proxy, requestBody.LinuxDoKey)
	m := logic.Trans(requestBody.Src, requestBody.Proxy, requestBody.LinuxDoKey)
	h := new(model.TranslateHistory)
	h.Src = requestBody.Src
	if found, err := h.FindBySrc(); err != nil {
		log.Printf("mysql查询发生错误:%+v\n", err)
	} else if found {
		log.Printf("在mysql中找到缓存%+v\n", h)
		rep := new(ResponseBody)
		rep.Dst = h.Dst
		rep.From = "cache"
		ctx.JSON(200, rep)
		return
	} else {
		log.Printf("没能在mysql中找到相同记录:%+v\n", h)
	}
	if dst, _, err := storage.GetTranslationFromRedis(requestBody.Src); err != nil {
		log.Printf("没能在redis中找到相同记录:%+v\n", h)
	} else {
		log.Printf("在redis中找到缓存%v\n", dst)
		rep := new(ResponseBody)
		rep.Dst = h.Dst
		rep.From = "cache"
		ctx.JSON(200, rep)
		return
	}
	var rep ResponseBody
	for k, v := range m {
		rep.From = k
		rep.Dst = v
		h := new(model.TranslateHistory)
		h.Src = requestBody.Src
		h.Dst = v
		h.From = k
		h.InsertOne()
		storage.InsertTranslationToRedis(requestBody.Src, v, k)
	}

	ctx.JSON(200, rep)
}
