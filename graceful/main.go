package main

import (
	"github.com/yadisnel/go-ms/v2/server"
	"github.com/yadisnel/go-ms/v2/server/mucp"
)

func main() {
	srv := mucp.NewServer(
		server.Wait(nil),
	)

	srv.Start()

	select {}
}
