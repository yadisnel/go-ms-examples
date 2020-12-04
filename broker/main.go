package main

import (
	"fmt"
	"log"
	"time"

	"github.com/yadisnel/go-ms/v2/broker"
	"github.com/yadisnel/go-ms/v2/broker/memory"
)

var (
	topic = "go.micro.topic.foo"
	b = memory.NewBroker()
)

func pub() {
	tick := time.NewTicker(time.Second)
	i := 0
	for _ = range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}
		if err := b.Publish(topic, msg); err != nil {
			log.Printf("[pub] failed: %v", err)
		} else {
			fmt.Println("[pub] pubbed message:", string(msg.Body))
		}
		i++
	}
}

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

	go pub()
	go sub()

	<-time.After(time.Second * 10)
}
