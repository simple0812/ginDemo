package router

import (
	"ginDemo/app"
	"ginDemo/controller"
)

func init() {
	ctrl := controller.AuthrCtrl{}

	app.Instance.GET("/authr/cookie", ctrl.GetCookie)
	app.Instance.POST("/authr/cookie", ctrl.SetCookie)
	app.Instance.GET("/authr/session", ctrl.GetSession)
	app.Instance.POST("/authr/session", ctrl.SetSession)
}
