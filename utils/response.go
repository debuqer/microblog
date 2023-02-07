package utils

import (
	"fmt"
	"net/http"
)

func Response(w http.ResponseWriter, res []byte) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(res))
}

type ErrorRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
