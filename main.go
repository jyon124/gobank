package main

import (
	"log"
)

func main() {
	store, err := NewPostgresstore()
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3001", store)
	server.Run()
}
