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
	"github.com/maha-1030/go-blog-api/handlers"
)

var (
	router     *mux.Router // router to serve the blog apis
	apiAddress string      // apiAddress stores the accessible address of blog api
)

const (
	DEFAULT_API_HOST = "localhost" // default host value at which api is accessible
	DEFAULT_API_PORT = "8080"      // default port value at which api is accessible
)

// main will start the blog api application
func main() {
	exportHealthCheckAPI()
	registerPublicHandlers()
	registerPrivateHandlers()
	startAPI()
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
	if host == "" {
		host = DEFAULT_API_HOST
	}

	port := os.Getenv("API_PORT")
	if port == "" {
		port = DEFAULT_API_PORT
	}

	apiAddress = host + ":" + port
}

// startAPI will start the api and starts listening to requests from apiAddress
func startAPI() {
	fmt.Println("Starting api at ", apiAddress+"/blog")
	err := http.ListenAndServe(apiAddress, router)
	if err != nil {
		fmt.Print(err)
	}
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")

		splittedBearerToken := strings.Split(bearerToken, " ")
		if len(splittedBearerToken) != 2 {
			handlers.RespondWithError(w, http.StatusUnauthorized, "Missing or invalid authorization token")

			return
		}

		if splittedBearerToken[0] != "Bearer" {
			handlers.RespondWithError(w, http.StatusUnauthorized, "Not a bearer token")

			return
		}

		token := splittedBearerToken[1]

		username, err := auth.ValidateToken(token)
		if err != nil {
			handlers.RespondWithError(w, http.StatusUnauthorized, "invalid authorization token")

			return
		}

		ctx := context.WithValue(r.Context(), "username", username)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func registerPublicHandlers() {
	authHandlers := handlers.GetAuthHandlers()

	router.HandleFunc("/register", authHandlers.Register).Methods(http.MethodPost)
	router.HandleFunc("/login", authHandlers.Login).Methods(http.MethodGet)
}

func registerPrivateHandlers() {
	privateRouter := router.NewRoute().Subrouter()

	privateRouter.Use(authMiddleware)
	fmt.Println("JWT authentication middleware is added")

	userHandlers := handlers.GetUserHandlers()

	privateRouter.HandleFunc("/user/{id}", userHandlers.Get).Methods(http.MethodGet)
	privateRouter.HandleFunc("/user/{id}", userHandlers.Update).Methods(http.MethodPut)
	privateRouter.HandleFunc("/user/{id}", userHandlers.Delete).Methods(http.MethodDelete)

	postHandlers := handlers.GetPostHandlers()

	privateRouter.HandleFunc("/post", postHandlers.Create).Methods(http.MethodPost)
	privateRouter.HandleFunc("/post/{id}", postHandlers.Get).Methods(http.MethodGet)
	privateRouter.HandleFunc("/post/{id}", postHandlers.Update).Methods(http.MethodPut)
	privateRouter.HandleFunc("/post/{id}", postHandlers.Delete).Methods(http.MethodDelete)

	commentHandlers := handlers.GetCommentHandlers()

	privateRouter.HandleFunc("/post/{postID}/comment", commentHandlers.Create).Methods(http.MethodPost)
	privateRouter.HandleFunc("/comment/{id}", commentHandlers.Get).Methods(http.MethodGet)
	privateRouter.HandleFunc("/post/{postID}/comment/{id}", commentHandlers.Update).Methods(http.MethodPut)
	privateRouter.HandleFunc("/comment/{id}", commentHandlers.Delete).Methods(http.MethodDelete)

	tagHandlers := handlers.GetTagHandlers()

	privateRouter.HandleFunc("/tag", tagHandlers.Create).Methods(http.MethodPost)
	privateRouter.HandleFunc("/tag/{id}", tagHandlers.Get).Methods(http.MethodGet)
	privateRouter.HandleFunc("/tag/{id}", tagHandlers.Update).Methods(http.MethodPut)
	privateRouter.HandleFunc("/tag/{id}", tagHandlers.Delete).Methods(http.MethodDelete)
}
