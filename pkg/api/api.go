package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "web/pkg/constant"
	"web/pkg/job"
)

func HelloWorld() func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.String(200, "Hello World")
	}
}

func StartJob() func(*gin.Context) {
	return func(c *gin.Context) {
		err := Server.Cron().AddJob("0 0/1 * * *", "startJob", job.Hello{})
		if err != nil {
			c.Status(http.StatusBadRequest)
			panic(err)
			return
		}

		c.String(http.StatusOK, "start")
	}
}

func RemoveJob() func(*gin.Context) {
	return func(c *gin.Context) {
		err := Server.Cron().RemoveJob("startJob")
		if err != nil {
			c.Status(http.StatusBadRequest)
			panic(err)
			return
		}
		c.String(http.StatusOK, "remove")
	}
}
