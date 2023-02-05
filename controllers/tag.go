package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/debuqer/microblog/models"
	"github.com/julienschmidt/httprouter"
)

func TagAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(r.Form.Get("q"))
	tags := models.GetAllTags(r.FormValue("q"))

	json, _ := json.Marshal(struct {
		Status string
		Data   []models.Tag
	}{
		"Success",
		tags,
	})
	fmt.Fprint(w, string(json))
}

func TagStore(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(w, ":)")
}

func TagShow(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(w, ":)")
}