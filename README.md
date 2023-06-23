
- [Rss Blog Aggregator](#rss-blog-aggregator)
  - [About the project](#about-the-project)
    - [API docs](#api-docs)
    - [Design](#design)
    - [Status](#status)
  - [Getting started](#getting-started)
    - [Layout](#layout)
  - [Notes](#notes)


# rss-blog-aggregator

## About the project

This is a simple backend project based on goLang. It is rss blog aggregator. That monitors registered blogs and fetched new blogs from them. Users can follow blogs and the blogs will be fetched from them. The server auto updates fetched blogs.

### API docs

API documentation is provided in the api_documentation.yml file.

### Status

The template project is in alpha status.

## Getting started

Below we'd like to describe the conventions or tools specific to Golang project.

### Layout

```tree
├── .gitignore
├── api_documentation.yml
├── go.mod
├── go.sum
├── handler-user.go
├── handler_feed.go
├── handler_feed_follow.go
├── main.go
├── middleware_auth.go
├── models.go
├── scraper.go
├── README.md
├── sqlc.go
├── sql
│   ├── queries
│   │   └── feed_follow.sql
│   │   └── feeds.sql
│   │   └── posts.sql
│   │   └── users.sql
│   └── schema
│       └── 001_user.sql
│       └── 002_user.sql
│       └── 003_feed.sql
│       └── 004_feed_follow.sql
│       └── 005_feed_lastfetch.sql
│       └── 006_posts.sql
└── vendor
    └── README.md


## Notes


