package main

import (
	"github.com/gin-gonic/gin"
	"rest-api/controllers"
)

func main(){
	router := gin.Default()
	v1 := router.Group("/api/v1/test")
	{
		v1.GET("", controllers.Test)
	}
	router.Run()
}
