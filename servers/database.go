package servers

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
	"log"
	"rest-api/models"
	"sync"
)

type database struct {
	db *gorp.DbMap
}

var instance *database
var once sync.Once

func Database() *database {
	once.Do(func() {
		db, err := sql.Open("mysql", "root:password@/testdb?parseTime=true")
		checkErr(err, "sql.Open failed")
		dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
		dbMap.AddTableWithName(models.Member{}, "member").SetKeys(true, "No")
		instance = &database{db: dbMap}
	})
	return instance
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func (db *database) GetInstance() *gorp.DbMap {
	return db.db
}
