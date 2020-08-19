package graph

import "fmt"

//BFSGraph graph
type BFSGraph struct {
	v             int // node counter
	queue         *Queue
	adjacencyList List
}

// NewBFSGraph constructor
func NewBFSGraph(i int) *BFSGraph {
	l := NewAdjacencyList(i)
	q := NewQueue()
	return &BFSGraph{
		v:             i,
		adjacencyList: l,
		queue:         &q,
	}
}

// Graph represents graph as a list of edges with nodes on the tips
type Graph [][2]int

// Len returns graph len
func (g Graph) Len() int {
	return len(g)
}

// NewGraph is a graph constuctor
func NewGraph(len int) Graph {
	return make(Graph, 0, len)
}

// AddEdge adds egde
func (g *Graph) AddEdge(a, b int) {
	*g = append(*g, [2]int{a, b})
}

// Unmarshal adds the graph from the appropriate data structure
func (b *BFSGraph) Unmarshal(s Graph) {
	for _, edge := range s {
		b.adjacencyList[edge[0]].Push(edge[1])
	}
}

// AddEdge adds
func (b *BFSGraph) AddEdge(s, d int) {
	b.adjacencyList[s].Push(d)
}

func (b *BFSGraph) convertToSlice(paths map[int]int, output []int, last int) []int {
	if val, ok := paths[last]; ok {
		output = append(output, val)
		return b.convertToSlice(paths, output, val)
	}
	return output
}

//FindShortestPath finds shortest path from src to dst and returns all nodes in between
//including src and dst
func (b *BFSGraph) FindShortestPath(s, dst int) ([]int, error) {
	if len(b.adjacencyList) < dst || len(b.adjacencyList) < s {
		return nil, fmt.Errorf("incorrect input parameters for the graph, src %d or dest %d are bigger then the graph len %d", s, dst, len(b.adjacencyList))
	}
	pathM, last := b.bfs(s, dst)
	if pathM == nil {
		return nil, nil
	}
	output := make([]int, 0, len(pathM))
	output = append(output, last)
	output = b.convertToSlice(pathM, output, last)
	// reverse slice
	outputRev := make([]int, 0, len(output))
	for i := len(output) - 1; i >= 0; i-- {
		outputRev = append(outputRev, output[i])
	}
	return outputRev, nil
}

// BFS main method
func (b *BFSGraph) bfs(s, dst int) (map[int]int, int) {
	pathM := make(map[int]int)
	visited := make(map[int]bool, b.v)
	visited[s] = true
	b.queue.Enqeue(s)
	for !b.queue.Empty() {
		s := b.queue.Deqeue()
		for _, i := range b.adjacencyList[s].GetConnectedNodes() {
			if !visited[i] {
				visited[i] = true
				b.queue.Enqeue(i)
				pathM[i] = s
				if i == dst {
					return pathM, i
				}
			}
		}
	}
	return nil, 0
}
