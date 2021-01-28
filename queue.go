package gokken

import "strings"

type Queue struct {
	head, tail int
	items      []string
}

func (q *Queue) Enqueue(item string) {
	q.items = append(q.items, item)
	q.tail++
}

func (q *Queue) Dequeue() string {
	q.head++
	return q.items[q.head-1]
}

func (q Queue) String() string {
	return strings.Join(q.items[q.head:q.tail], ", ")
}
