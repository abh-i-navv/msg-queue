package main

import (
	"time"

	"github.com/abh-i-navv/msg-queue/internal/queue"
)

func main() {
	broker := queue.NewBroker((100))

	go producer(broker)
	time.Sleep(1 * time.Second)
	go consumer(broker)

	time.Sleep(1 * time.Second)
}
