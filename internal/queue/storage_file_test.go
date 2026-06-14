package queue

import (
	"os"
	"testing"
)

func TestFileStorage_SaveLoad(t *testing.T) {
	path := "test.log"
	defer os.Remove(path)

	storage := NewFileStorage(path)

	msg := Message{
		ID:      "1",
		Payload: []byte("hello"),
	}

	if err := storage.Save(msg); err != nil {
		t.Fatalf("save failde: %v", err)
	}

	messages, err := storage.Load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}

	if len(messages) != 1 {
		t.Fatalf("expected 1 message, got %d", len(messages))
	}

	if messages[0].ID != "1" {
		t.Fatalf("expected id=1 got %s", messages[0].ID)
	}

	if string(messages[0].Payload) != "hello" {
		t.Fatalf(
			"expected payload hello got %s",
			string(messages[0].Payload),
		)
	}

}
