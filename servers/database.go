package servers

import (
	"database/sql"
	"gopkg.in/gorp.v1"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Database struct {
	db *gorp.DbMap
}

var Db *Database
var database = GetInstance()

func GetInstance() *Database {
	if Db == nil {
		db, err := sql.Open("mysql", "root:password@/testdb")
		checkErr(err, "sql.Open failed")
		dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
		Db = &Database{db : dbmap}
	}
	return Db
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Select() *Database {

}
