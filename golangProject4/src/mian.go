package main

import "github.com/gin-gonic/gin"
import ""

func main() {
	r := GetRouter()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务

}
