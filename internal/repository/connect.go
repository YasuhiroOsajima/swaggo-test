package repository

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Connect("mysql", "root:root@127.0.0.1/practice")
	if err != nil {
		log.Fatalf("DB connect failed. err: %s", err)
		panic(err)
	}
}
