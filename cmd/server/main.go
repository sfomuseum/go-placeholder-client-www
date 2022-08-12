package main

import (
	"context"
	"github.com/sfomuseum/go-placeholder-client-www/application/server"
	"log"
)

func main() {

	ctx := context.Background()
	err := server.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run server, %v", err)
	}
}
