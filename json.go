package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with error 5XX: ", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	RespondWithJson(w, code, errResponse{Error: msg})

}
func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, error := json.Marshal(payload)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
	//json.NewEncoder(w).Encode(payload)
}
