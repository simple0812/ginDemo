package controller

import (
	"encoding/json"
	"fmt"
	"ginDemo/model"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FooCtrl struct{}

func (FooCtrl) Get(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "hello %s", name)
}

func (FooCtrl) Render(c *gin.Context) {
	c.HTML(http.StatusOK, "foo.html", gin.H{
		"key": "Main website",
	})
}

func (FooCtrl) Json(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user": struct {
		Name   string
		Gender int
	}{Name: "a", Gender: 1}})
}

func (FooCtrl) Post(c *gin.Context) {
	fmt.Println(c.Request.Body)
	_ = "breakpoint"

	result, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	fmt.Println(result)
	var user model.User
	json.Unmarshal(result, &user)
	fmt.Println(user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Printfx(args ...interface{}) {
	for _, val := range args {
		fmt.Println(val)
	}
}
