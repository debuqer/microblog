package models

import (
	"log"

	"github.com/debuqer/microblog/utils"
)

type Event struct {
	Id          int
	Label       string
	Description string
	DateTime    ComplexDate
	Timeline    Timeline
}

func (e *Event) Save() (*Event, error) {
	db, err := utils.DB()
	if err != nil {
		log.Fatal(err)
		return nil, error(err)
	}
	defer db.Close()

	if e.Id != 0 {
		stmt, err := db.Prepare("INSERT INTO events(Label, Description, Era, Year, Month, Day, timeline_id) VALUES(?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			return nil, error(err)
		}

		affected, _ := stmt.Exec(e.Label, e.Description, e.DateTime.Era, e.DateTime.Date.Year, e.DateTime.Date.Month, e.DateTime.Date.Day, e.Timeline.Id)
		id, _ := affected.LastInsertId()
		e.Id = int(id)
	} else {
		stmt, err := db.Prepare("UPDATE events SET Label = ?, Description = ?, Era = ?, Year = ?, Month = ?, Day = ?, timeline_id = ? WHERE Id = ?")
		if err != nil {
			return nil, error(err)
		}

		stmt.Exec(e.Label, e.Description, e.DateTime.Era, e.DateTime.Date.Year, e.DateTime.Date.Month, e.DateTime.Date.Day, e.Timeline.Id, e.Id)
	}

	return e, nil
}
