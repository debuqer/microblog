package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/debuqer/microblog/models"
	"github.com/julienschmidt/httprouter"
)

func TagAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tags, err := models.GetAllTags(r.FormValue("q"))

	var res []byte
	if err != nil {
		res, _ = json.Marshal(struct {
			Status  string
			message string
		}{
			"Fail",
			"Failed to get tags",
		})
	} else {
		res, _ = json.Marshal(struct {
			Status string
			Data   []models.Tag
		}{
			"Success",
			tags,
		})
	}

	fmt.Fprint(w, string(res))
}

func TagStore(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	createdBy, _ := strconv.Atoi(r.PostFormValue("created_by"))
	t := models.Tag{
		Label:       r.PostFormValue("label"),
		Description: r.PostFormValue("description"),
		Color:       r.PostFormValue("color"),
		CreatedBy:   createdBy,
		CreatedDate: r.PostFormValue("created_date"),
	}

	t.Save()

	res, _ := json.Marshal(struct {
		Status string
		Data   models.Tag
	}{
		"Success",
		t,
	})
	fmt.Fprint(w, res)
}

func TagShow(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(w, ":)")
}
