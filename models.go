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

type feedFollow struct {
	ID        uuid.UUID `json:"id"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func feedFollowToFeedFollow(dbFeedFollow *db.FeedFollow) feedFollow {
	return feedFollow{
		ID:        dbFeedFollow.ID,
		FeedID:    dbFeedFollow.FeedID,
		UserID:    dbFeedFollow.UserID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
	}
}

func feedFollowsArrayToFeedFollowsArray(dbFeedFollow []db.FeedFollow) []feedFollow {
	feedFollows := make([]feedFollow, len(dbFeedFollow))
	for i, f := range dbFeedFollow {
		feedFollows[i] = feedFollowToFeedFollow(&f)
	}
	return feedFollows
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	FeedID      uuid.UUID `json:"feed_id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Url         string    `json:"url"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func postToPost(dbPost *db.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		FeedID:      dbPost.FeedID,
		Title:       dbPost.Title,
		Description: description,
		Url:         dbPost.Url,
		PublishedAt: dbPost.PublishedAt,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
	}
}

func postArrayToPostArray(dbPost []db.Post) []Post {
	posts := make([]Post, len(dbPost))
	for i, p := range dbPost {
		posts[i] = postToPost(&p)
	}
	return posts
}
