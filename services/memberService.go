package services

import (
	"rest-api/models"
	"rest-api/repositories"
	"time"
)

type MemberService struct{}

var memberRepository repositories.MemberRepository

func (self MemberService) GetMember() [] models.Member {
	return memberRepository.GetMember()
}

func newMember(name, id string) models.Member {
	return models.Member{
		Name:   name,
		RegDate: time.Now(),
		UpdDate: time.Now(),
		Id:    id,
	}
}

func (self MemberService) CreateMember(name, id string) models.Member {
	member := newMember(name, id)
	return memberRepository.CreateMember(member)
}