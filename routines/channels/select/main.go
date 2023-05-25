package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	var i int64 = 0

	go func() {
		for {
			time.Sleep(1 * time.Second)
			atomic.AddInt64(&i, 1)
			msg := Message{i, "RabbitMQ"}
			c1 <- msg
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Kafka"}
			c2 <- msg
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Printf("Message %d from %s\n", msg.id, msg.Msg)

		case msg := <-c2:
			fmt.Printf("Message %d from %s\n", msg.id, msg.Msg)

		case <-time.After(3 * time.Second):
			fmt.Println("Timeout")
		}
	}
}
