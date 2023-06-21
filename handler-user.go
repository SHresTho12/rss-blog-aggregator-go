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

	RespondWithJson(w, 201, UserToUser(&user))
}

func (apiConfig *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user db.User) {
	// apikey, err := auth.GetApi(r.Header)
	// if err != nil {
	// 	RespondWithError(w, 403, "User not created")
	// 	return
	// }
	// user, err := apiConfig.DB.GetUserByApiKey(r.Context(), apikey)
	// if err != nil {
	// 	RespondWithError(w, 404, "User not found")
	// 	return
	// }

	RespondWithJson(w, 200, UserToUser(&user))

}

func (apiConfig *apiConfig) handlerUserGetPosts(w http.ResponseWriter, r *http.Request, user db.User) {

	posts, err := apiConfig.DB.GetPostForUser(r.Context(), db.GetPostForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		RespondWithError(w, 500, "Could not get posts")
		return
	}

	defer r.Body.Close()

	RespondWithJson(w, 201, postArrayToPostArray(posts))

}
