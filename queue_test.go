package gokken

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := Queue{}
	queue.Enqueue("first")
	queue.Enqueue("second")
	queue.Enqueue("third")

	if queue.Dequeue() != "first" {
		t.Error("Didn't remove from the beginning!")
	}
	if fmt.Sprint(queue) != "second, third" {
		t.Error("Didn't remove from the beginning!")
	}

	if queue.Dequeue() != "second" {
		t.Error("Didn't remove from the beginning!")
	}
	if fmt.Sprint(queue) != "third" {
		t.Error("Didn't remove from the beginning!")
	}
}
