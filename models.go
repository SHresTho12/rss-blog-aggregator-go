package main

import (
	"time"

	"github.com/google/uuid"
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

type feed struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Url    string    `json:"url"`
	UserID uuid.UUID `json:"user_id"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}

func feedToFeed(dbFeed *db.Feed) feed {
	return feed{
		ID:     dbFeed.ID,
		Name:   dbFeed.Name,
		Url:    dbFeed.Url,
		UserID: dbFeed.UserID,
		// CreatedAt: dbFeed.CreatedAt,
		// UpdatedAt: dbFeed.UpdatedAt,
	}
}

func feedsArrayToFeedsArray(dbFeed []db.Feed) []feed {
	feeds := make([]feed, len(dbFeed))
	for i, f := range dbFeed {
		feeds[i] = feedToFeed(&f)
	}
	return feeds
}
