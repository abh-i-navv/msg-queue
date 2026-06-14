package queue

type Storage interface {
	Save(Message) error
	Load() ([]Message, error)
}
