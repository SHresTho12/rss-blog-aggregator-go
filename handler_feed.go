package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/shrestho12/go-practice-project/internal/db"
)

func (apiConfig *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user db.User) {

	type param struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	p := param{}
	err := decoder.Decode(&p)
	if err != nil {
		RespondWithError(w, 400, "Invalid request payload")
		return
	}

	feed, err := apiConfig.DB.CreateFeed(r.Context(), db.CreateFeedParams{
		ID:        uuid.New(),
		Name:      p.Name,
		Url:       p.Url,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
	})
	if err != nil {
		RespondWithError(w, 500, "Feed not created")
		return
	}

	defer r.Body.Close()

	RespondWithJson(w, 201, feedToFeed(&feed))

}
