package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/debuqer/microblog/models"
	"github.com/debuqer/microblog/utils"
	"github.com/julienschmidt/httprouter"
)

func TagAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tags, err := models.GetAllTags(r.FormValue("q"))

	var res []byte
	if err != nil {
		res, _ = json.Marshal(utils.ErrorRes{
			Status:  "Fail",
			Message: "Failed to get tags",
		})
	} else {
		res, _ = json.Marshal(struct {
			Status string       `json:"status"`
			Data   []models.Tag `json:"data"`
		}{
			"Success",
			tags,
		})
	}

	fmt.Fprint(w, string(res))
}

func TagStore(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var res []byte

	decoder := json.NewDecoder(r.Body)
	t := models.Tag{}
	err := decoder.Decode(&t)
	if err != nil {
		res, _ = json.Marshal(utils.ErrorRes{
			Status:  "Fail",
			Message: "Provide Data as json in request",
		})
	} else {
		t.Save()

		res, _ = json.Marshal(struct {
			Status string     `json:"status"`
			Data   models.Tag `json:"data"`
		}{
			"Success",
			t,
		})
	}

	fmt.Fprint(w, string(res))
}

func TagShow(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(w, ":)")
}
