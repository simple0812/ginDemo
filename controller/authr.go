package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthrCtrl struct{}

func (AuthrCtrl) GetCookie(c *gin.Context) {
	if cookie, err := c.Request.Cookie("test"); err == nil {
		c.String(http.StatusOK, cookie.Value)
		return
	}

	c.String(http.StatusOK, "cookie Name is empty")
}

func (AuthrCtrl) SetCookie(c *gin.Context) {
	//当前url才有的cookie
	cookie := &http.Cookie{
		Name:  "test",
		Value: "0",
	}
	//golang cookie跟url有直接关系
	cookie1 := &http.Cookie{
		Name:  "zl",
		Value: "1",
		Path:  "/", //全局cookie
	}

	http.SetCookie(c.Writer, cookie)
	http.SetCookie(c.Writer, cookie1)
	c.String(http.StatusOK, "cookie set success")
}

func (AuthrCtrl) GetSession(c *gin.Context) {
	c.String(http.StatusOK, "hello %s", "world")
}

func (AuthrCtrl) SetSession(c *gin.Context) {
	c.String(http.StatusOK, "hello %s", "world")
}
