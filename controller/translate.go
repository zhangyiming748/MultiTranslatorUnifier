package controller

import (
	"fmt"
	"log"

	"github.com/zhangyiming748/MultiTranslatorUnifier/logic"

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
	Src  string `json:"src"`
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
	s := new(storage.SQLiteStorage)
	var rep ResponseBody
	if dst, err := s.GetTranslation(requestBody.Src); err != nil {
		log.Printf("查询数据库失败,err = %v\n", err)
	} else if dst == "" {
		log.Println("数据库查询为空")
		m := logic.Trans(requestBody.Src, requestBody.Proxy, requestBody.LinuxDoKey)
		for k, v := range m {
			rep.From = k
			rep.Dst = v
			rep.Src = requestBody.Src
		}
		s.SaveTranslation(rep.From, rep.Src, rep.Dst)
	} else {
		rep.From = "sqlite"
		rep.Dst = dst
		rep.Src = requestBody.Src
	}
	ctx.JSON(200, rep)
}
