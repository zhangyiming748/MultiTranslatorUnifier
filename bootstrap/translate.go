package bootstrap

import (
	"github.com/zhangyiming748/MultiTranslatorUnifier/controller"

	"github.com/gin-gonic/gin"
)

func InitTranslate(engine *gin.Engine) {
	routeGroup := engine.Group("/api")
	{
		c := new(controller.TranslateController)
		routeGroup.GET("/v1/translate", c.GetTranslate)
		routeGroup.POST("/v1/translate", c.PostTranslate)
	}
}
