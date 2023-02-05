package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	fmt.Println(w, ":)")
}

func TagShow(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(w, ":)")
}
