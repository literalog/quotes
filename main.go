package main

import (
	"log"
)

func main() {
	store, err := NewMongoStorage()
	if err != nil {
		log.Fatal(err)
	}

	api := NewApiServer(":8080", store)
	api.Run()
}
