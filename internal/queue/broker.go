package queue

import (
	"fmt"
	"sync"
	"time"
)

type InflightMessage struct {
	Message     Message
	DeliveredAt time.Time
}

type Broker struct {
	messages chan Message
	storage  Storage
	done     chan struct{}

	inflight map[string]InflightMessage
	mu       sync.Mutex

	visibilityTimeout time.Duration
}

func NewBroker(bufferSize int, storage Storage) (*Broker, error) {

	b := &Broker{
		messages:          make(chan Message, bufferSize),
		storage:           storage,
		done:              make(chan struct{}),
		inflight:          make(map[string]InflightMessage),
		visibilityTimeout: 10 * time.Second,
	}

	messages, err := storage.Load()
	if err != nil {
		return nil, err
	}

	for _, msg := range messages {
		b.messages <- msg
	}
	return b, nil
}

func (b *Broker) Publish(msg Message) error {

	if err := b.storage.Save(msg); err != nil {
		return err
	}

	b.messages <- msg
	return nil
}

func (b *Broker) Consume() (*Delivery, bool) {
	msg, ok := <-b.messages

	if !ok {
		return nil, false
	}

	b.mu.Lock()

	b.inflight[msg.ID] = InflightMessage{
		Message:     msg,
		DeliveredAt: time.Now(),
	}

	b.mu.Unlock()

	return &Delivery{
		Message: msg,
		ack: func() error {
			return b.Ack(msg.ID)
		},
	}, true
}

func (b *Broker) Close() {
	close(b.done)
	close(b.messages)
}

func (b *Broker) Ack(id string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	delete(b.inflight, id)

	return nil
}

func (b *Broker) retryExpired() {
	b.mu.Lock()
	defer b.mu.Unlock()

	for id, msg := range b.inflight {
		msg.Message.RetryCount++

		fmt.Printf(
			"Retrying %s (%d)\n",
			msg.Message.ID,
			msg.Message.RetryCount,
		)
		if time.Since(msg.DeliveredAt) >
			b.visibilityTimeout {

			delete(b.inflight, id)

			b.messages <- msg.Message
		}
	}
}

func (b *Broker) StartRetryLoop() {
	ticker := time.NewTicker(time.Second)

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				b.retryExpired()
			case <-b.done:
				return
			}
		}
	}()
}
