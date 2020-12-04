package main

import (
	"context"

	"github.com/yadisnel/go-ms/v2/service"
	"github.com/yadisnel/go-ms/v2/service/mucp"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, name *string, msg *string) error {
	*msg = "Hello " + *name
	return nil
}

func main() {
	// create new service
	service := mucp.NewService(
		service.Name("greeter"),
	)

	// set the handler
	service.Server().Handle(
		service.Server().NewHandler(new(Greeter)),
	)

	// run service
	service.Run()
}
