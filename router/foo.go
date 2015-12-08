package router

import (
	"fmt"
	"ginDemo/app"
	"ginDemo/controller"

	"github.com/gin-gonic/gin"
)

func init() {
	ctrl := controller.FooCtrl{}
	app.Instance.GET("/foo/view", func(c *gin.Context) {
		fmt.Println("/foo/view")
		c.Next()
		fmt.Println("/foo/view end")
	}, ctrl.Render)

	app.Instance.GET("/foo/json", ctrl.Json)
	app.Instance.POST("/foo/json", ctrl.Post)
	app.Instance.GET("/name/:name", ctrl.Get)
}
