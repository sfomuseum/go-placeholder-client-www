package main

import (
	"context"
	"log"

	"github.com/sfomuseum/go-placeholder-client-www/application/server"
)

func main() {

	ctx := context.Background()
	logger := log.Default()

	err := server.Run(ctx, logger)

	if err != nil {
		logger.Fatalf("Failed to run server, %v", err)
	}
}
