package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var router = gin.Default()

func GetRouter()  *gin.Engine{

	index()
	user()
	return router
}

func index()  {
	index := router.Group("/")
	// 添加 Get 请求路由
	index.GET("/", retHelloGinAndMethod)
	// 添加 Post 请求路由
	index.POST("/", retHelloGinAndMethod)
	// 添加 Put 请求路由
	index.PUT("/", retHelloGinAndMethod)
	// 添加 Delete 请求路由
	index.DELETE("/", retHelloGinAndMethod)
	// 添加 Patch 请求路由
	index.PATCH("/", retHelloGinAndMethod)
	// 添加 Head 请求路由
	index.HEAD("/", retHelloGinAndMethod)
	// 添加 Options 请求路由
	index.OPTIONS("/", retHelloGinAndMethod)
}
func user()  {
	user := router.Group("/user")
	// 添加 Get 请求路由
	user.GET("/:name", retUserGinAndMethod)
	// 添加 Post 请求路由
	user.POST("/:name", retUserGinAndMethod)
	// 添加 Put 请求路由
	user.PUT("/:name", retUserGinAndMethod)
	// 添加 Delete 请求路由
	user.DELETE("/:name", retUserGinAndMethod)
	// 添加 Patch 请求路由
	user.PATCH("/:name", retUserGinAndMethod)
	// 添加 Head 请求路由
	user.HEAD("/:name", retUserGinAndMethod)
	// 添加 Options 请求路由
	user.OPTIONS("/:name", retUserGinAndMethod)
}

func retHelloGinAndMethod(context *gin.Context) {
	//context.String(http.StatusOK, "hello gin "+strings.ToLower(context.Request.Method)+" method")
	context.JSON(http.StatusOK, gin.H{
		"message": "hello gin "+strings.ToLower(context.Request.Method)+" method",
	})
}
func retUserGinAndMethod(context *gin.Context)  {
	username := context.Param("name")
	context.JSON(http.StatusOK, gin.H{
		"message": "用户："+ username +"已经保存成功" + strings.ToLower(context.Request.Method)+" method",
	})
}
