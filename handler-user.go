package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/shrestho12/go-practice-project/internal/db"
)

func (apiConfig *apiConfig) handleUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ready")

	type param struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	p := param{}
	err := decoder.Decode(&p)
	if err != nil {
		RespondWithError(w, 400, "Invalid request payload")
		return
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), db.CreateUserParams{
		ID:        uuid.New(),
		Name:      p.Name,
		Email:     p.Email,
		Password:  p.Password,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		RespondWithError(w, 500, "User not created")
		return
	}

	defer r.Body.Close()

	RespondWithJson(w, 200, UserToUser(&user))
}
