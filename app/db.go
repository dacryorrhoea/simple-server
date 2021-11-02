package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func initDB() (err error) {
	var key = "root:Password1@tcp(localhost:3306)/usersdb"

	Db, err = sqlx.Connect("mysql", key)
	if err != nil {
		return
	}

	err = Db.Ping()
	return
}
