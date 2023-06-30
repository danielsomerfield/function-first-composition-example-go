package main

import (
	"function-first-composition-example-go/review-server/configuration"
	"function-first-composition-example-go/review-server/server"
	"log"
)

func main() {
	reviewServer := server.NewServer(configuration.FromEnv)

	if err := reviewServer.Start(); err != nil {
		log.Fatalf("failed to start server: %s", err.Error())
	}

	<-reviewServer.Shutdown
}
