package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/models"
	"rest-api/services"
)

type MemberController struct{}

var Member MemberController

var memberService services.MemberService

func (self MemberController) GetMember(c *gin.Context) {
	memberService.GetMember()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusCreated, "result": "success", "data": memberService.GetMember()})
}

func (self MemberController) CreateMember(c *gin.Context) {
	var member models.Member;
	c.ShouldBindJSON(&member);
	result := memberService.CreateMember(member.Name, member.Id)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "result": "success", "data": result})
}
