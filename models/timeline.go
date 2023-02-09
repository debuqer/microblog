package models

import (
	"log"

	"github.com/debuqer/microblog/utils"
)

type Timeline struct {
	Id           int
	Slug         string
	Title        string
	Descriptiuon string
	CreatedBy    int
	CreatedDate  string
}

func (t *Timeline) Save() (rett *Timeline, reterr error) {
	db, err := utils.DB()
	if err != nil {
		log.Fatal(err)
		return nil, error(err)
	}
	defer db.Close()

	if t.Id == 0 {
		stmt, err := db.Prepare("INSERT INTO timelines(slug, title, description, created_by, created_at) VALUES(?, ?, ?, ?, ?)")
		reterr = err

		affected, _ := stmt.Exec(t.Slug, t.Title, t.Descriptiuon, t.CreatedBy, t.CreatedDate)
		id, _ := affected.LastInsertId()
		t.Id = int(id)
	} else {
		stmt, err := db.Prepare("UPDATE timelines SET slug = ?, title = ?, description = ?, updated_by = ?, updated_at = ? WHERE id = ?")
		reterr = err

		stmt.Exec(t.Slug, t.Title, t.Descriptiuon, t.CreatedBy, t.CreatedDate, t.Id)
	}

	rett = t
	return
}
