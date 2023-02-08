package main

import "github.com/gin-gonic/gin"

func hello(c *gin.Context) {
	c.String(200, "helloworld")
}

func main() {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.GET("/", hello)
	router.Run(":8080")
}
