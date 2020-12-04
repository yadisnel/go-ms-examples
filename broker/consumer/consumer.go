package main

import (
	"fmt"
	"log"

	"github.com/yadisnel/go-ms/v2/broker"
	"github.com/yadisnel/go-ms/v2/broker/memory"
)

var (
	topic = "go.micro.topic.foo"
	b     = memory.NewBroker()
)

// Example of a shared subscription which receives a subset of messages
func sharedSub() {
	_, err := b.Subscribe(topic, func(m *broker.Message) error {
		fmt.Println("[sub] received message:", string(m.Body), "header", m.Header)
		return nil
	}, broker.Queue("consumer"))
	if err != nil {
		fmt.Println(err)
	}
}

// Example of a subscription which receives all the messages
func sub() {
	_, err := b.Subscribe(topic, func(m *broker.Message) error {
		fmt.Println("[sub] received message:", string(m.Body), "header", m.Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	if err := b.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := b.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	sub()
	select {}
}
