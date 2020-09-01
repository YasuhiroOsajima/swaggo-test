package repository

import (
	"log"
	"os"
	"time"

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

	for i := 0; i < 5; i++ {
		db, err = sqlx.Connect("mysql", user+":"+password+"@tcp("+server+":"+port+")/practice")
		if err != nil {
			log.Fatalf("DB connect failed. Retry after 10 sec. err: %s", err)
			time.Sleep(time.Second * 10)
			continue
		}
	}
	if err != nil {
		log.Fatalf("DB connect failed. err: %s", err)
		panic(err)
	}
}
