package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/z876730060/ToTServer/pkg"
	api "github.com/z876730060/web/pkg/api"
	. "github.com/z876730060/web/pkg/constant"
	"net/http"
)

func Setup() {
	Server = pkg.NewServer().InitLog("web").InitCron().SetBanner("banner.txt")
	e := gin.New()
	e.Use(gin.Recovery(), gin.Logger())
	// 加载路由
	api.InitRoute(e)

	webServer := &http.Server{
		Addr:    ":8080",
		Handler: e.Handler(),
	}

	Server.LoadWebService(webServer)
	_ = Server.Start()
}
