package stack

import (
	"sync"
)

// NewStack represents a stack constructor
func NewStack() Stack {
	return Stack{
		data: make([]interface{}, 0, 100),
	}
}

// Stack represents data type stack
type Stack struct {
	mu   sync.RWMutex
	data []interface{}
}

// Push push data to stack
func (s *Stack) Push(i interface{}) {
	s.mu.Lock()
	s.data = append(s.data, i)
	s.mu.Unlock()
}

// Pop pops data out of stack
func (s *Stack) Pop() interface{} {
	s.mu.RLock()
	ln := len(s.data) - 1
	data := s.data[ln]
	s.mu.RUnlock()
	return data
}

//Len shows the lenghth of the stack
func (s *Stack) Len() int {
	s.mu.RLock()
	ln := len(s.data)
	s.mu.RUnlock()
	return ln
}

//IsEmpty returns if the stack is empty
func (s *Stack) IsEmpty() bool {
	s.mu.RLock()
	ln := len(s.data)
	s.mu.RUnlock()
	return ln == 0
}
