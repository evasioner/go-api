package models

import "time"

type Member struct {
	No      int64  `db:"member_no"`
	Name    string `db:"member_name"`
	Id      string `db:"member_id"`
	RegDate time.Time `db:"member_reg_date"`
	UpdDate time.Time `db:"member_upd_date"`
}
