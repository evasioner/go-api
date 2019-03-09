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

	member := router.Group("/api/v1/member")
	{
		member.GET("", controllers.Member.GetMember)
	}
	router.Run()
}

