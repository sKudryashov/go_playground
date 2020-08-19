package graph

import (
	"sync"
)

// NewStack represents a stack constructor
func NewStack() Stack {
	return Stack{
		data: make([]string, 0, 100),
	}
}

// Stack represents data type stack
type Stack struct {
	mu   sync.RWMutex
	data []string
}

// Push push data to stack
func (s *Stack) Push(i string) {
	s.mu.Lock()
	s.data = append(s.data, i)
	s.mu.Unlock()
}

func (s *Stack) remove(i int) {
	if i == 0 {
		s.data = s.data[:0]
		return
	}
	if i == len(s.data)-1 {
		s.data = s.data[i:]
		return
	}
	left := s.data[:i-1]
	right := s.data[i+1:]
	left = append(left, right...)
	s.data = left
}

// Remove removes given ID from the stack
func (s *Stack) Remove(ID string) {
	for i := 0; i < len(s.data); i++ {
		if s.data[i] == ID {
			s.remove(i)
			break
		}
	}
}

// Pop pops data out of stack
func (s *Stack) Pop() (string, bool) {
	if len(s.data) == 0 {
		return "", false
	}
	ln := len(s.data) - 1
	data := s.data[ln]
	s.data = s.data[:ln]
	return data, true
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
