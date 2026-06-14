package queue

import (
	"bufio"
	"encoding/json"
	"os"
)

type FileStorage struct {
	path string
}

func NewFileStorage(path string) *FileStorage {
	return &FileStorage{
		path: path,
	}
}

func (s *FileStorage) Save(msg Message) error {
	file, err := os.OpenFile(
		s.path,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)

	if err != nil {
		return err
	}

	defer file.Close()

	data, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	_, err = file.Write(append(data, '\n'))

	return err
}

func (s *FileStorage) Load() ([]Message, error) {
	file, err := os.Open(s.path)
	if os.IsNotExist(err) {
		return []Message{}, nil
	}
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var messages []Message

	for scanner.Scan() {
		var msg Message

		if err := json.Unmarshal(scanner.Bytes(), &msg); err != nil {
			return nil, err
		}

		messages = append(messages, msg)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
