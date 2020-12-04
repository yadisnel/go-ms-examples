package main

import (
	"fmt"

	"github.com/yadisnel/go-ms/v2/config"
	"github.com/yadisnel/go-ms/v2/config/memory"
	"github.com/yadisnel/go-ms/v2/config/source/file"
)

func main() {
	// load the config from a file source
	c, err := memory.NewConfig(config.WithSource(
		file.NewSource(file.WithPath("./config.json")),
	))
	if err != nil {
		fmt.Println(err)
		return
	}

	// define our own host type
	type Host struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	}

	var host Host

	// read a database host
	v, err := c.Load("hosts", "database")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := v.Scan(&host); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(host.Address, host.Port)
}
