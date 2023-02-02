package routes

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ApiRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", hello)

	return router
}

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
