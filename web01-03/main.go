package main

import (
	"github.com/gin-gonic/gin"
)

func HelloGin(context *gin.Context){
	context.JSON(200,gin.H{
		"message" : "Hello Gin",
	})
}

func main(){
	r := gin.Default()
	r.GET("/book", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message" : "HelloGet",
		})
	})
	r.POST("/book", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message" : "HelloPost",
		})
	})
	r.PUT("/book", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message" : "HelloPut",
		})
	})
	r.DELETE("/book", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message" : "HelloDelete",
		})
	})

	r.Run()
}