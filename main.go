package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var router *mux.Router

func main() {
	exportHealthCheckAPI()

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	fmt.Println("host: ", host, " --- port: ", port)

	err := http.ListenAndServe(host+":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}

func init() {
	//Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error occured while loading .env file, err: %v\n", err)

		return
	}

	router = mux.NewRouter().Headers("Content-Type", "application/json").Subrouter()
}

func exportHealthCheckAPI() {
	router.HandleFunc("/.well-known/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	}).Methods(http.MethodGet)
}
