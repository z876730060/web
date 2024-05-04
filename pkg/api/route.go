package api

import (
	"github.com/gin-gonic/gin"
)

func InitRoute(e *gin.Engine) {
	e.GET("/", HelloWorld())
	e.GET("/start", StartJob())
	e.GET("/remove", RemoveJob())
}
