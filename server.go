package main

import (
	"ginDemo/app"
	_ "ginDemo/router"
)

func main() {
	router := app.Instance

	router.Run(":3005")
}
