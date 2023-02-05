package models

import (
	"database/sql"
	"log"

	"github.com/debuqer/microblog/utils"
)

type Tag struct {
	Id          int
	Label       string
	Color       string
	Description string
	CreatedBy   User
	CreatedDate string
}

func GetAllTags(q string) []Tag {
	db := utils.DB()
	defer db.Close()

	var res *sql.Rows
	var err error
	var stmt *sql.Stmt

	if q == "" {
		stmt, err = db.Prepare("SELECT * FROM tags LIMIT 10")
		if err != nil {
			log.Fatal(err)
		}

		res, err = stmt.Query()
	} else {
		stmt, err = db.Prepare("SELECT * FROM tags WHERE label like ? LIMIT 10")
		if err != nil {
			log.Fatal(err)
		}

		res, err = stmt.Query("%" + q + "%")
	}
	if err != nil {
		log.Fatal(err)
	}

	tags := make([]Tag, 0)
	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var t Tag
		res.Scan(&t.Id, &t.Label, &t.Color, &t.Description, &t.CreatedBy, &t.CreatedDate)
		tags = append(tags, t)
	}

	return tags
}
