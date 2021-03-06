package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// --- Configuration --- //

var config Configuration

func getConfig(ENV string) Configuration {
	file, err := os.Open(fmt.Sprintf("config.%s.json", ENV))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var config Configuration
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func main() {
	// Get configuration
	ENV := os.Getenv("ENV")
	if ENV == "" {
		ENV = "dev"
	}
	fmt.Println(fmt.Sprintf("Running in ENV: %s", ENV))
	config = getConfig(ENV)

	db = connectDb(config.Db)
	defer db.Close(context.Background())
	pingDb(db)

	r := mux.NewRouter()

	r.HandleFunc("/publish/{uuid}", publish)
	r.HandleFunc("/publish_done/{uuid}", publishDone)

	r.HandleFunc("/admin/clear_streams/{user_id}", clearStreams)

	port := 9000
	fmt.Println(fmt.Sprintf("Serving on port %d", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
