package controllers

import (
	"github.com/gin-gonic/gin"
	"rest-api/services"
)

type TestController struct {
}

var Service services.TestService
var Tests TestController

func (self TestController) Test(c *gin.Context){

	c.JSON(200, gin.H{
		"message": Service.Test(),
	})
}
