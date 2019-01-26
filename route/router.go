package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"rest-api/controllers"
	"time"
)

func Route() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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

