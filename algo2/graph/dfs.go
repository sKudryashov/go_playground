package graph

// DFSGraph d
type DFSGraph struct {
	v             int // node counter
	queue         *Queue
	adjacencyList List
}

// NewDFSGraph constructor
func NewDFSGraph(v int) *DFSGraph {
	l := NewAdjacencyList(v)
	q := NewQueue()
	return &DFSGraph{
		v:             v,
		adjacencyList: l,
		queue:         &q,
	}
}

//AddEdge adds edge
func (b *DFSGraph) AddEdge(s, d int) {
	b.adjacencyList[s].Push(d)
}

func (b *DFSGraph) dFSUtil(v int, visited []bool) {
	visited[v] = true
	for i := b.adjacencyList[v].Begin(); i <= b.adjacencyList[v].End(); i++ {
		if !visited[i] {
			b.dFSUtil(v, visited)
		}
	}
}

//DFS launch loop
func (b *DFSGraph) DFS(i int) {
	visited := make([]bool, b.v, b.v)
	b.dFSUtil(i, visited)

}

// LaunchDFS laucnher dfs
func LaunchDFS() {
	g := NewDFSGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)
	g.DFS(2)
}
