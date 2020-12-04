package main

import (
	"github.com/yadisnel/go-ms-examples/helloworld/handler"
	"github.com/yadisnel/go-ms/v2/logger"
	"github.com/yadisnel/go-ms/v2/service"
	"github.com/yadisnel/go-ms/v2/service/mucp"

	pb "github.com/yadisnel/go-ms-examples/helloworld/proto"
)

func main() {
	// New Service
	helloworld := mucp.NewService(
		service.Name("helloworld"),
		service.Version("latest"),
	)

	// Initialise service
	helloworld.Init()

	// Register Handler
	pb.RegisterHelloworldHandler(helloworld.Server(), new(handler.Helloworld))

	// Run service
	if err := helloworld.Run(); err != nil {
		logger.Fatal(err)
	}
}
