package main

import (
	"log"

	"github.com/prathameshj610/fampay-youtube-assignment/internal/database"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/server"
)

func main() {

	db, err := database.NewDatabaseClient()

	if err != nil {
		log.Fatalf("failed to initialize db client: %s", err)
	}

	srv := server.NewEchoServer(db)

	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}

}
