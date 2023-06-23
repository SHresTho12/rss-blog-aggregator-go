package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println("Hello World")
	portStr := os.Getenv("PORT")

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

	newRouter := chi.NewRouter()
	newRouter.Get("/ready", handleReady)
	newRouter.Get("/error", handleError)
	router.Mount("/api", newRouter)

	fmt.Println("Port: ", portStr)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("Port: ", portStr)
}
