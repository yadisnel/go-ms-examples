package main

import (
	"context"
	"fmt"

	"github.com/yadisnel/go-ms/v2/client"
	"github.com/yadisnel/go-ms/v2/client/mucp"
)

func main() {
	c := mucp.NewClient()

	request := c.NewRequest("greeter", "Greeter.Hello", "john", client.WithContentType("application/json"))
	var response string

	if err := c.Call(context.TODO(), request, &response); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}
