package model

import (
	"github.com/sKudryashov/graphsrv/pkg/graph"
	"github.com/sKudryashov/graphsrv/pkg/pb"
)

//GraphModel is a data model i.e. wrapper over any transport layer and responsible for data conversion/filtering/validation
// between business logic/storage and transport layer
type GraphModel struct {
	graph *pb.Graph
}

// SearchResult represents graph search interaction
type SearchResult struct {
	ID    string
	Nodes []int
}

// SearchModel represents graph search interaction
type SearchModel struct {
	graph *pb.Search
}

//NewSearch constructor for seach model
func NewSearch(graph *pb.Search) SearchModel {
	return SearchModel{
		graph: graph,
	}
}

//GetSrcDst returns a pair of nodes in between of which we are going to find the shortest path
func (sm SearchModel) GetSrcDst() (int, int) {
	return int(sm.graph.Src), int(sm.graph.Dst)
}

//NewGraphModel is a wrapper over appropriate GRPC api entity
func NewGraphModel(graph *pb.Graph) GraphModel {
	return GraphModel{
		graph: graph,
	}
}

// GetGraph returns given graph representation
func (g GraphModel) GetGraph() graph.Graph {
	edges := graph.NewGraph(len(g.graph.Edges))
	for _, e := range g.graph.Edges {
		edges.AddEdge(int(e.A), int(e.B))
	}
	return edges
}
