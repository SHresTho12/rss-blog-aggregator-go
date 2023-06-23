package main

import (
	"fmt"
	"net/http"
)

func handleReady(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ready")
	RespondWithJson(w, 200, map[string]string{"message": "ready"})
}
