package queue

import (
	"testing"
)

func TestQueueNewQueue(t *testing.T) {
	q := NewQueue[int](10)
	if !q.Empty() {
		t.Errorf("Expected queue to be empty, it was not")
	}

	if q.Full() {
		t.Errorf("Expected queue to be empty, but it was full")
	}
}

func TestQueueEnqueue(t *testing.T) {
	q := NewQueue[int](1)
	q.Enqueue(1)

	if err := q.Enqueue(1); err == nil {
		t.Errorf("Expected enqueue on full queue to return error, got %+v", err)
	}

	if item, ok := q.Peek(); ok == false || item != 1 {
		t.Errorf("Expected first item to be 1, it was %d and ok was %v", item, ok)
	}

	if !q.Full() {
		t.Errorf("Expected queue to be full")
	}

	if err := q.Enqueue(2); err == nil {
		t.Errorf("Expected error when attempting to enqueue on full queue, got nil")
	}
}

func TestQueueDequeue(t *testing.T) {
	testItem := 1
	q := NewQueue[int](1)
	q.Enqueue(testItem)

	if item, err := q.Dequeue(); err != nil || item != testItem {
		t.Errorf("Expected item to be 1 and no error, got %d and %+v", item, err)
	}

	if !q.Empty() {
		t.Errorf("Expected queue to be empty")
	}

	if _, err := q.Dequeue(); err == nil {
		t.Errorf("Expected error when attempting to dequeue on empty queue, was <nil>")
	}
}

