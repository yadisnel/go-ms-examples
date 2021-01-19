package main

import (
	"net/http"

	"github.com/yadisnel/go-ms/v2examples/template/web/handler"
	"github.com/yadisnel/go-ms/v2/util/log"
	"github.com/yadisnel/go-ms/v2/web"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.template"),
		web.Version("latest"),
	)

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/example/call", handler.ExampleCall)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
