package main

import (
	"log"

	"github.com/ishanshre/gotodo/pkg/middleware"
	"github.com/ishanshre/gotodo/pkg/storage"
)

func main() {
	store, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatalf("error in connecting to database: %v", err)
	}
	if err := store.Init(); err != nil {
		log.Fatalf("Error in creating table : %v", err)
	}
	server := middleware.NewAPIServer(":8000", store)
	server.Run()
}
