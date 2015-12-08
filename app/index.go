package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var Instance *gin.Engine

func init() {
	Instance = gin.Default()

	Instance.Static("/public", "./public")
	Instance.LoadHTMLGlob("views/*")

	Instance.Use(func(c *gin.Context) {
		fmt.Println("middleware start")
		c.Next()
		fmt.Println("middleware end")
	})
}
