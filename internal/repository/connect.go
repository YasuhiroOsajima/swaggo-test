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
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASS")

	var err error
	db, err = sqlx.Connect("mysql", user+":"+password+"@tcp("+server+":"+port+")/practice")
	if err != nil {
		log.Fatalf("DB connect failed. err: %s", err)
		panic(err)
	}
}
