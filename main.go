package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", hello)

	http.ListenAndServe(":8080", router)
}
