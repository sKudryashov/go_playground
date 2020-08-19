package graph_test

import (
	"testing"

	"golang.org/x/exp/errors/fmt"

	"github.com/sKudryashov/algo/graph"
)

func TestBFS(t *testing.T) {
	// graph.LaunchBFS()
	g := graph.NewBFSGraph(8)
	// TODO: where to fill in the queue ?
	// g.AddEdge(0, 1)
	// g.AddEdge(0, 2)
	// g.AddEdge(1, 2)
	// // g.AddEdge(2, 0)
	// g.AddEdge(2, 3)
	// // g.AddEdge(3, 3)
	// g.AddEdge(2, 3)
	// g.AddEdge(2, 3)
	g.AddEdge(1, 2)
	g.AddEdge(2, 4)
	g.AddEdge(1, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(4, 6)
	g.AddEdge(2, 5)
	path := g.FindShortestPath(2, 5)
	fmt.Println(" shortest path ", path)
	// g.BFS(2)
}

func TestDeptFirst(t *testing.T) {
	graph.LaunchDFS()
}
