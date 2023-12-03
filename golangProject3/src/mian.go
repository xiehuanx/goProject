package main


import "golangProject3/src/router"

func main() {
	r := router.GetRouter()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务

}
