package main

import (
	"fmt"
	"time"

	"github.com/abh-i-navv/msg-queue/internal/queue"
)

func main() {
	storage := queue.NewFileStorage("./data")
	broker, err := queue.NewBroker(100, storage)

	if err != nil {
		fmt.Println("Error loading messages")
		return
	}
	broker.StartRetryLoop()

	go producer(broker)
	go consumer(broker)

	time.Sleep(10 * time.Second)
}
