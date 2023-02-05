package utils

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DB() *sql.DB {
	sql, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/old-old")
	if err != nil {
		log.Fatal(err)
	}

	return sql
}
