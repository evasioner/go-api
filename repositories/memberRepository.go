package repositories

import "rest-api/servers"

type MemberRepository struct {

}

var Member MemberRepository

func GetMember() {
	db := servers.GetInstance()
	db
}