package main

import (
	"time"

	"github.com/shrestho12/go-practice-project/internal/db"
)

type user struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	ApiKey    string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func UserToUser(dbUser *db.User) user {
	return user{
		ID:        dbUser.ID.String(),
		Name:      dbUser.Name,
		Email:     dbUser.Email,
		ApiKey:    dbUser.ApiKey,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}
