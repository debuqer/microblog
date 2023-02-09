package models

import (
	"database/sql"
	"log"

	"github.com/debuqer/microblog/utils"
)

type Tag struct {
	Id          int    `json:"id,omitempty"`
	Label       string `json:"label"`
	Color       string `json:"color"`
	Description string `json:"description"`
	CreatedBy   int    `json:"created_by",omitempty`
	CreatedDate string `json:"created_date"`
}

func (t *Tag) Save() (*Tag, error) {
	db, err := utils.DB()
	if err != nil {
		log.Fatal(err)
		return nil, error(err)
	}
	defer db.Close()

	if t.Id == 0 {
		stmt, err := db.Prepare("INSERT INTO tags(Label, Description, Color, created_by, created_at) VALUES(?, ?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
			return nil, error(err)
		}

		affected, _ := stmt.Exec(t.Label, t.Description, t.Color, t.CreatedBy, t.CreatedDate)
		id, _ := affected.LastInsertId()
		t.Id = int(id)
	} else {
		stmt, err := db.Prepare("UPDATE tags SET Label = ?, Description = ?, Color = ?, updated_by = ?, updated_at = ? WHERE Id = ?")
		if err != nil {
			log.Fatal(err)
			return nil, error(err)
		}

		stmt.Exec(t.Label, t.Description, t.Color, t.CreatedBy, t.CreatedDate, t.Id)
	}

	return t, nil
}

func (t *Tag) Fetch() error {
	db, err := utils.DB()
	if err != nil {
		log.Fatal(err)
		return error(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM tags WHERE id = ? LIMIt 1")
	if err != nil {
		log.Fatal(err)
		return error(err)
	}

	stmt.QueryRow(t.Id).Scan(&t.Id, &t.Label, &t.Color, &t.Description, &t.CreatedBy, &t.CreatedDate)

	return nil
}

func GetAllTags(q string) ([]Tag, error) {
	var res *sql.Rows
	var err error
	var stmt *sql.Stmt

	db, err := utils.DB()
	if err != nil {
		return nil, error(err)
	}
	defer db.Close()

	if q == "" {
		stmt, err = db.Prepare("SELECT * FROM tags LIMIT 10")
		if err != nil {
			return nil, error(err)
		}

		res, err = stmt.Query()
	} else {
		stmt, err = db.Prepare("SELECT * FROM tags WHERE label like ? LIMIT 10")
		if err != nil {
			return nil, error(err)
		}

		res, err = stmt.Query("%" + q + "%")
	}
	if err != nil {
		log.Fatal(err)
		return nil, error(err)
	}

	tags := make([]Tag, 0)

	for res.Next() {
		var t Tag
		res.Scan(&t.Id, &t.Label, &t.Color, &t.Description, &t.CreatedBy, &t.CreatedDate)
		tags = append(tags, t)
	}

	return tags, nil
}
