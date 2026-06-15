package queue

type Delivery struct {
	Message Message

	ack func() error
}

func (d *Delivery) Ack() error {
	return d.ack()
}
