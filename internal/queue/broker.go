package queue

type Broker struct {
	messages chan Message
	done     chan struct{}
}

func NewBroker(bufferSize int) *Broker {
	return &Broker{
		messages: make(chan Message, bufferSize),
	}
}

func (b *Broker) Publish(msg Message) {
	b.messages <- msg
}

func (b *Broker) Consume() (Message, bool) {
	msg, ok := <-b.messages

	return msg, ok
}

func (b *Broker) Close() {
	close(b.done)
	close(b.messages)
}
