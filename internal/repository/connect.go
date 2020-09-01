package repository

import (
	"log"
	"os"
	"strings"
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
	waitMsg := "connect: connection refused"
	notInclude := -1

	for i := 0; i < 5; i++ {
		db, err = sqlx.Connect("mysql", user+":"+password+"@tcp("+server+":"+port+")/practice")
		if strings.Index(err.Error(), waitMsg) != notInclude {
			log.Fatalf("DB connect failed. err: %s", err)
			time.Sleep(time.Second * 10)
			continue
		}
		break
	}
	if err != nil {
		log.Fatalf("DB connect failed. err: %s", err)
		panic(err)
	}
}
