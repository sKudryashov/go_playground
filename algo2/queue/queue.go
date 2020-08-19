package queue

import (
	"sync"
)

//NewFixedQueue represents queue constructor
func NewFixedQueue(len int) *Queue {
	return &Queue{
		data: make([]interface{}, 0, len),
	}
}

// Queue represents unlimited queeue data structure
type Queue struct {
	mu   sync.RWMutex
	data []interface{}
}

// Push pushes data into the end of the queue
func (q *Queue) Push(item interface{}) {
	q.mu.Lock()
	q.data = append(q.data, item)
	q.mu.Unlock()
}

// Pop pops data from the end of the queue
func (q *Queue) Pop() interface{} {
	q.mu.Lock()
	ln := len(q.data)
	data := q.data[ln-1]
	q.data = q.data[0 : ln-2]
	q.mu.Unlock()
	return data
}

// Deqeue dequeues data from the beginning of the queue
func (q *Queue) Deqeue() interface{} {
	var first interface{}
	q.mu.RLock()
	ln := len(q.data)
	first = q.data[0]
	q.mu.RUnlock()
	if ln > 1 {
		q.mu.Lock()
		q.data = q.data[1 : len(q.data)-1]
		q.mu.Unlock()
	} else {
		q.mu.Lock()
		q.data = q.data[:0]
		q.mu.Unlock()
	}
	return first
}

//Enqeue pushes data to the beginning of the queue
func (q *Queue) Enqeue(item interface{}) {
	q.mu.Lock()
	tmp := q.data
	data := make([]interface{}, 1, len(q.data)+1)
	data[0] = item
	data = append(data, tmp...)
	q.mu.Unlock()
}

// Len queue len
func (q *Queue) Len() int {
	q.mu.RLock()
	ln := len(q.data)
	q.mu.RUnlock()
	return ln

}

// Empty Len queue len
func (q *Queue) Empty() bool {
	q.mu.RLock()
	e := (len(q.data) == 0)
	q.mu.RUnlock()
	return e
}
