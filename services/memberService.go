package services

import (
	"rest-api/models"
	"rest-api/repositories"
)

type MemberService struct {}

var Member repositories.MemberRepository

func (self MemberService) GetMember() [] models.Member{
	return Member.GetMember()
}