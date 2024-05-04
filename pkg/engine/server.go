package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/z876730060/ToTServer/pkg"
	"github.com/z876730060/web/docs"
	"github.com/z876730060/web/pkg/api"
	. "github.com/z876730060/web/pkg/constant"
)

func Setup() {
	initSwagger()

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

func initSwagger() {
	docs.SwaggerInfo.Title = "ToT Server API"
	docs.SwaggerInfo.Description = "开发调试平台"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}
}
