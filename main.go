package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/maha-1030/go-blog-api/auth"
)

var (
	router     *mux.Router // router to serve the blog apis
	apiAddress string      // apiAddress stores the accessible address of blog api
)

// main will start the blog api application
func main() {
	exportHealthCheckAPI()

	err := http.ListenAndServe(apiAddress, router)
	if err != nil {
		fmt.Print(err)
	}
}

// init to load configs and creates router and apiAddress
func init() {
	// load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error occured while loading .env file, err: %v\n", err)

		return
	}

	// intializes the router
	router = mux.NewRouter().Headers("Content-Type", "application/json").PathPrefix("/blog").Subrouter()

	setApiAddress()
}

// setApiAddress will set the api address where api is accessible
func setApiAddress() {
	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")

	apiAddress = host + ":" + port
}

// exportHealthCheckAPI will add the public health check endpoint to router
func exportHealthCheckAPI() {
	router.HandleFunc("/.well-known/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	}).Methods(http.MethodGet)
}

// authMiddleware authenticates the requests by validating bearer token
func authMiddleware(next http.Handler) http.Handler {
	defer fmt.Println("JWT authentication middleware is added")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")

		splittedBearerToken := strings.Split(bearerToken, " ")
		if len(splittedBearerToken) != 2 {
			http.Error(w, "Missing or invalid authorization token", http.StatusUnauthorized)

			return
		}

		if splittedBearerToken[0] != "Bearer" {
			http.Error(w, "Not a bearer token", http.StatusUnauthorized)

			return
		}

		token := splittedBearerToken[1]

		username, err := auth.ValidateToken(token)
		if err != nil {
			http.Error(w, "invalid authorization token", http.StatusUnauthorized)

			return
		}

		ctx := context.WithValue(r.Context(), "username", username)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
