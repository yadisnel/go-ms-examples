package main

import (
	"github.com/yadisnel/go-ms/v2examples/template/srv/handler"
	"github.com/yadisnel/go-ms/v2examples/template/srv/subscriber"
	"github.com/yadisnel/go-ms/v2"
	"github.com/yadisnel/go-ms/v2/util/log"

	example "github.com/yadisnel/go-ms/v2examples/template/srv/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.template"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.template", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.template", service.Server(), subscriber.Handler)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
