package queue

type Broker struct {
	messages chan Message
	storage  Storage
	done     chan struct{}
}

func NewBroker(bufferSize int, storage Storage) (*Broker, error) {

	b := &Broker{
		messages: make(chan Message, bufferSize),
		storage:  storage,
		done:     make(chan struct{}),
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

func (b *Broker) Consume() (Message, bool) {
	msg, ok := <-b.messages

	return msg, ok
}

func (b *Broker) Close() {
	close(b.done)
	close(b.messages)
}
