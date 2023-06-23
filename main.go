package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	//"github.com/flowchartsman/swaggerui"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/shrestho12/go-practice-project/internal/db"
)

type apiConfig struct {
	DB *db.Queries
}

// @Version 1.0.0
// @Title RSS Aggregation
// @Description A RSS aggregation that follows specified blogs and scrapes their feeds
// @ContactEmail abce@email.com
// @Server http://localhost:8000/api/ Server-1

func main() {

	godotenv.Load()
	fmt.Println("Hello World")
	portStr := os.Getenv("PORT")
	dbString := os.Getenv("DB_URL")
	if dbString == "" {
		log.Fatal("DB_URL must be set")
	}

	conn, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal("Can not connect to DB", err)
	}

	apiCon := apiConfig{
		DB: db.New(conn),
	}

	go startScraping(apiCon.DB, 10, time.Minute)

	//create a new router
	router := chi.NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portStr,
	}

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// spec, err := ioutil.ReadFile("./oas.yml")
	// if err != nil {
	// 	fmt.Println("Error")
	// }

	// http.Handle("/swagger/", http.StripPrefix("/swagger", swaggerui.Handler(spec)))

	newRouter := chi.NewRouter()
	newRouter.Get("/ready", handleReady)
	newRouter.Get("/error", handleError)
	//newRouter.Get("/user", apiCon.handleUser)
	newRouter.Post("/user", apiCon.handleCreateUser)
	newRouter.Get("/user", apiCon.authMiddleware(apiCon.handleGetUser))
	newRouter.Post("/feed", apiCon.authMiddleware(apiCon.handleCreateFeed))
	newRouter.Get("/feed", apiCon.handlerGetFeed)
	newRouter.Post("/feedFollow", apiCon.authMiddleware(apiCon.handleCreateFeedFollow))
	newRouter.Get("/feedFollow", apiCon.authMiddleware(apiCon.handlerGetFeedFollow))
	newRouter.Delete("/feedFollow/{feedFollowId}", apiCon.authMiddleware(apiCon.handlerDeleteFeedFollow))
	newRouter.Get("/posts", apiCon.authMiddleware(apiCon.handlerUserGetPosts))
	//newRouter.Get("/swagger", handleSwagger)
	router.Mount("/api", newRouter)

	fmt.Println("Port: ", portStr)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("Port: ", portStr)
}
