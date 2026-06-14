package main

import (
	"fmt"
	"time"

	"github.com/abh-i-navv/msg-queue/internal/queue"
)

func producer(b *queue.Broker) {
	for i := 1; i <= 10; i++ {
		msg := queue.Message{
			ID:      fmt.Sprintf("%d", i),
			Payload: []byte(fmt.Sprintf("message-%d", i)),
		}
		fmt.Println("Published:", string(msg.Payload))

		b.Publish(msg)

		time.Sleep(time.Second)
	}

}
