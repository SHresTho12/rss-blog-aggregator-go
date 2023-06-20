package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/shrestho12/go-practice-project/internal/db"
)

func (apiConfig *apiConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {

	type param struct {
		Feed_id uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	p := param{}
	err := decoder.Decode(&p)
	if err != nil {
		RespondWithError(w, 400, "Invalid request payload")
		return
	}

	feedFollow, err := apiConfig.DB.CreateFeedFollow(r.Context(), db.CreateFeedFollowParams{

		ID:        uuid.New(),
		FeedID:    p.Feed_id,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		RespondWithError(w, 500, "Feed not created")
		return
	}

	defer r.Body.Close()

	RespondWithJson(w, 201, feedFollowToFeedFollow(&feedFollow))

}

func (apiConfig *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {

	feedFollows, err := apiConfig.DB.GetFeedFollow(r.Context(), user.ID)
	if err != nil {
		RespondWithError(w, 500, "Could not get feed")
		return
	}

	defer r.Body.Close()

	RespondWithJson(w, 201, feedFollowsArrayToFeedFollowsArray(feedFollows))

}

func (apiConfig *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {

	feedFollowIdstr := chi.URLParam(r, "feedFollowId")
	feedFollowId, err := uuid.Parse(feedFollowIdstr)
	if err != nil {
		RespondWithError(w, 400, "Invalid request payload")
		return
	}

	err = apiConfig.DB.DeleteFeedFollow(r.Context(), db.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		RespondWithError(w, 500, "Could not delete feed")
		return
	}

	defer r.Body.Close()

	RespondWithJson(w, 201, "Feed deleted")

}
