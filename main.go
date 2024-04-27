package main

import (
	"github.com/gin-gonic/gin"
	"github.com/z876730060/ToT/pkg"
	"net/http"
)

func main() {
	server := pkg.NewServer().InitLog("web")
	e := gin.New()
	e.Use(gin.Recovery(), gin.Logger())
	e.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!!!")
	})

	webServer := &http.Server{
		Addr:    ":8080",
		Handler: e.Handler(),
	}

	server.LoadWebService(webServer)
	_ = server.Start()
}
