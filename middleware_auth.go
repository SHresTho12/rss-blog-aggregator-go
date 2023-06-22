package main

import (
	"net/http"

	"github.com/shrestho12/go-practice-project/internal/auth"
	"github.com/shrestho12/go-practice-project/internal/db"
)

type authHandler func(http.ResponseWriter, *http.Request, db.User)

func (apiConfig *apiConfig) authMiddleware(next authHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetApi(r.Header)
		if err != nil {
			RespondWithError(w, 403, err.Error())
			return
		}
		user, err := apiConfig.DB.GetUserByApiKey(r.Context(), apikey)
		if err != nil {
			RespondWithError(w, 404, "User not found")
			return
		}
		next(w, r, user)
	})
}
