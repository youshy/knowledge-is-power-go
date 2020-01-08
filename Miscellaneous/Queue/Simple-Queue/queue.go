package queue

import (
	"sync"
)

// This is a generic type
type Item int

// ItemQueue is the queue of all items
type ItemQueue struct {
	items []Item
	lock  sync.RWMutex
}

// New creates a new ItemQueue
func (q *ItemQueue) New() *ItemQueue {
	q.items = []Item{}
	return q
}

// Enqueue adds an Item to the end of the queue
func (q *ItemQueue) Enqueue(i Item) {
	q.lock.Lock()
	q.items = append(q.items, i)
	q.lock.Unlock()
}

// Dequeue removes an Item from the start of the queue
func (q *ItemQueue) Dequeue() *Item {
	q.lock.Lock()
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	q.lock.Unlock()
	return &item
}

// Front returs the Item next in the queue, without removing it
func (q *ItemQueue) Front() *Item {
	q.lock.RLock()
	item := q.items[0]
	q.lock.RUnlock()
	return &item
}

// IsEmpty returns true if the queue is empty
func (q *ItemQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of Items in the queue
func (q *ItemQueue) Size() int {
	return len(q.items)
}
