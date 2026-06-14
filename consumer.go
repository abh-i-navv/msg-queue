package main

import (
	"fmt"

	"github.com/abh-i-navv/msg-queue/internal/queue"
)

func consumer(b *queue.Broker) {
	for {
		msg, ok := b.Consume()
		if !ok {
			fmt.Println("Not OK")
			return
		}
		fmt.Println("Consumed:", string(msg.Payload))
	}
}
