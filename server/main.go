package main

import (
	"log"

	"github.com/yadisnel/go-ms-examples/server/handler"
	"github.com/yadisnel/go-ms/v2/server"
	"github.com/yadisnel/go-ms/v2/server/mucp"
)

func main() {
	// Initialise Server
	server := mucp.NewServer(
		server.Name("go.micro.srv.example"),
	)

	// Register Handlers
	server.Handle(
		server.NewHandler(
			new(handler.Example),
		),
	)

	// Run server
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	select {}
}
