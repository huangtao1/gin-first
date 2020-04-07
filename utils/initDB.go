package utils

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "")
	if err != nil {
		log.Panic("connect to db failed:" + err.Error())
	}
}
