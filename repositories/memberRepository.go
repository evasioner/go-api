package repositories

import (
	"rest-api/models"
	"rest-api/servers"
)

type MemberRepository struct{}

func (self MemberRepository) GetMember() [] models.Member {
	db := servers.Database().GetInstance()
	var members []models.Member
	db.Select(&members, "select * from member")
	return members
}
