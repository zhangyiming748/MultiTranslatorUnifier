package controller

import (
	"MultiTranslatorUnifier/logic"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TranslateController struct{}

/*
 */
func (t TranslateController) GetTranslate(ctx *gin.Context) {
	user := ctx.Query("user")
	ctx.String(200, fmt.Sprintf("%s还活着!", user))
}

// 结构体必须大写 否则找不到
type RequestBody struct {
	Src        string `json:"src"`
	Proxy      string `json:"dst,omitempty"`
	LinuxDoKey string `json:"linuxdokey,omitempty`
}

type ResponseBody struct {
	From string `json:"from"`
	Dst  string `json:"dst"`
}

/*
 */
func (t TranslateController) PostTranslate(ctx *gin.Context) {
	fmt.Println("接收到post请求")
	var requestBody RequestBody
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		fmt.Println(requestBody)
	}
	fmt.Println(requestBody.Src, requestBody.Proxy)
	m := logic.Trans(requestBody.Src, requestBody.Proxy, requestBody.LinuxDoKey)
	var rep ResponseBody
	for k, v := range m {
		rep.From = k
		rep.Dst = v
	}
	ctx.JSON(200, rep)
}
