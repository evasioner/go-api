package controllers

import (
	"github.com/gin-gonic/gin"
	"rest-api/services"
)

type MemberController struct {}

var Member MemberController

var memberService services.MemberService

func (self MemberController) GetMember(c *gin.Context){
	c.JSON(200, gin.H{
		"data": memberService.GetMember(),
	})
}
