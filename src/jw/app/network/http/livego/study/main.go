package main

import (
	//"github.com/gwuhaolin/livego"
	"github.com/gin-gonic/gin"
)


func main() {
	e := gin.Engine{}
	e.GET("/", func(c *gin.Context) {

	})

	e.Run()
}
