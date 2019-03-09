package models

type Member struct {
	No      int64  `db:"member_no"`
	Name    string `db:"member_name"`
	Id      string `db:"member_id"`
	RegDate string `db:"member_reg_date"`
	UpdDate string `db:"member_upd_date"`
}
