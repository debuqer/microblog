package handlers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func TimelineAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(w, ":)")
}

func TimelineStore(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(w, ":)")
}

func TimelineShow(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(w, ":)")
}

func TimelineUpdate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(w, ":)")
}

func TimelineClone(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(w, ":)")
}

func TimelineMerge(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(w, ":)")
}

func TimelineReaction(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(w, ":)")
}

func TimelineEvents(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(w, ":)")
}

func TimelineLatest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(w, ":)")
}
