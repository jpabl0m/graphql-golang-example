package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose"
)

const dir = "./db/migrations"

func main() {
	driver := "mysql"
	if err := goose.SetDialect(driver); err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open(driver, fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASS"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DBNAME"),
	))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Fatal(goose.Run("up", db, dir))
}
