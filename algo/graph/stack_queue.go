package graph

// NewStack represents a stack constructor
func NewStack() Stack {
	return Stack{
		data: make([]int, 0, 100),
	}
}

// Stack represents data type stack
type Stack struct {
	data []int
}

// Push push data to stack
func (s *Stack) Push(i int) {
	s.data = append(s.data, i)
}

// Pop pops data out of stack
func (s Stack) Pop() int {
	return s.data[len(s.data)-1]
}

//NewQueue represents queue constructor
func NewQueue() Queue {
	return Queue{
		data: make([]int, 0, 100),
	}
}

// Queue represents unlimited queeue data structure
type Queue struct {
	data []int
}

// Deqeue push data from queue
func (q *Queue) Deqeue() int {
	var first int
	ln := len(q.data)
	if ln > 1 {
		first = q.data[0]
		q.data = q.data[1:]
	} else {
		first = q.data[0]
		q.data = q.data[:0]

	}
	return first
}

//Enqeue e
func (q *Queue) Enqeue(item int) {
	q.data = append(q.data, item)
}

// Len queue len
func (q *Queue) Len() int {
	return len(q.data)
}

// Empty Len queue len
func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

type List map[int]*VertexList

// VertexList represents vertexes list
type VertexList []int

func (v *VertexList) Push(i int) {
	*v = append(*v, i)
}

func (v VertexList) Begin() int {
	return v[0]
}

func (v VertexList) End() int {
	end := v[len(v)-1]
	return end
}

func (v VertexList) GetConnectedNodes() []int {
	return v
}

// NewAdjacencyList constructor for adjlist
func NewAdjacencyList(i int) List {
	asjList := make(List, i)
	for j := 0; j < i; j++ {
		vlist := make(VertexList, 0, i)
		asjList[j] = &vlist
	}
	return asjList
}
