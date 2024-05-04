package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoute(e *gin.Engine) {
	v1 := e.Group("/v1")
	{
		v1.GET("/", HelloWorld())
		v1.GET("/start", StartJob())
		v1.GET("/remove", RemoveJob())
	}

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
