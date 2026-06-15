package main

import (
	"fmt"

	"github.com/abh-i-navv/msg-queue/internal/queue"
)

func consumer(b *queue.Broker) {
	for {
		delivery, ok := b.Consume()
		if ok {
			fmt.Println("Consumed:", string(delivery.Message.ID))

			if err := delivery.Ack(); err != nil {
				fmt.Println(err)
			}
		}
	}
}
