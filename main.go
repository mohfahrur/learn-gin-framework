package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.Llongfile)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello world")
	})
	r.Run()
}
