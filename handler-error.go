package main

import (
	"fmt"
	"net/http"
)

func handleError(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ready")
	RespondWithError(w, 400, "Error while responding")
}
