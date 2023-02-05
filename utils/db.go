package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func db() *sql.DB {
	sql, err := sql.Open("mysql", "root:@127.0.0.1:3306/old-old")
	if err != nil {
		panic(err)
	}

	defer sql.Close()

	return sql
}
