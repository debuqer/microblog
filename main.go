package main

import (
	"net/http"

	"github.com/debuqer/microblog/routes"
)

func main() {
	apiRouter := routes.ApiRouter()

	http.ListenAndServe(":8083", apiRouter)
}
