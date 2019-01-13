package route

import (
	"github.com/gin-gonic/gin"
	"rest-api/controllers"
)

func Route(){
	router := gin.Default()
	test := router.Group("/api/v1/test")
	{
		test.GET("", controllers.Tests.Test)
	}

	member := router.Group("/api/v1/new")
	{
		member.GET("", controllers.News.News)
	}
	router.Run()
}
