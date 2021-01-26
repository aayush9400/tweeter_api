package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/anujc4/tweeter_api/handler"
	"github.com/anujc4/tweeter_api/internal/env"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func routeHandler(handler *handler.HttpApp) {
	router := mux.NewRouter().StrictSlash(true)

	apiV1 := router.PathPrefix("/v1").Subrouter()

	// Health Check Endpoint
	router.HandleFunc("/simple_health", handler.SimpleHealthCheck).Methods("GET")
	router.HandleFunc("/detail_health", handler.DetailedHealthCheck).Methods("GET")

	// Users API
	apiV1.HandleFunc("/user", handler.CreateUser).Methods("POST")
	apiV1.HandleFunc("/users", handler.GetUsers).Methods("GET")
	apiV1.HandleFunc("/user/{user_id}", handler.GetUserByID).Methods("GET")
	apiV1.HandleFunc("/user/{user_id}", handler.UpdateUser).Methods("PUT")
	apiV1.HandleFunc("/user/{user_id}", handler.DeleteUser).Methods("DELETE")

	// Tweets
	// TODO: Implement the handler
	// apiV1.HandleFunc("/tweet", handler.CreateTweet).Methods("POST")
	// apiV1.HandleFunc("/tweets", handler.GetTweets).Methods("GET")
	// apiV1.HandleFunc("/tweet/{tweet_id}", handler.GetTweetByID).Methods("GET")
	// apiV1.HandleFunc("/tweet/{tweet_id}", handler.UpdateTweet).Methods("PUT")
	// apiV1.HandleFunc("/tweet/{tweet_id}", handler.UpdateTweet).Methods("DELETE")

	// Start the server
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(":3000", handlers.RecoveryHandler()(loggedRouter)))
}

func main() {
	env := env.Init()
	h := handler.NewHttpApp(env.DB)

	fmt.Println("Starting Tweeter API on port 3000...")

	routeHandler(h)
}
