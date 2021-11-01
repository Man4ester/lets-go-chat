package handlers

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	response := "Hello world"
	fmt.Fprint(w, response)
}
