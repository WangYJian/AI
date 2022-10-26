package router

import (
	"AI/router/api"
	"github.com/gin-gonic/gin"
)

func UseMyRouter(f *gin.Engine) {
	g := f.Group("/get")
	{
		//上传操作信息并返回图片
		g.GET("/generate/:text/:style", api.Generate)
	}

}