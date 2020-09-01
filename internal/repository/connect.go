package repository

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	server := os.Getenv("DBSERVER")
	port := os.Getenv("DBPORT")

	var err error
	db, err = sqlx.Connect("mysql", "root:root@"+server+":"+port+"/practice")
	if err != nil {
		log.Fatalf("DB connect failed. err: %s", err)
		panic(err)
	}
}
