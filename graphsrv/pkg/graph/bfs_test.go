package graph

import (
	"reflect"
	"testing"
)

func TestBFSGraph_BFS(t *testing.T) {
	type args struct {
		src   int
		dst   int
		graph [][2]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			want: []int{2, 5},
			args: args{
				src: 2,
				dst: 5,
				graph: [][2]int{
					{1, 2},
					{2, 4},
					{1, 3},
					{3, 4},
					{4, 5},
					{4, 6},
					{2, 5},
				},
			},
		},
		{
			want: []int{2, 4, 5},
			args: args{
				src: 2,
				dst: 5,
				graph: [][2]int{
					{1, 2},
					{2, 4},
					{1, 3},
					{3, 4},
					{4, 5},
					{4, 6},
				},
			},
		},
		{
			want: []int{1, 3, 6},
			args: args{
				src: 1,
				dst: 6,
				graph: [][2]int{
					{1, 2},
					{2, 4},
					{1, 3},
					{3, 4},
					{4, 5},
					{4, 6},
					{3, 6},
				},
			},
		},
		{
			want: []int{1, 3, 6, 8},
			args: args{
				src: 1,
				dst: 8,
				graph: [][2]int{
					{1, 2},
					{2, 4},
					{1, 3},
					{3, 4},
					{4, 5},
					{4, 6},
					{3, 6},
					{6, 8},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBFSGraph(len(tt.args.graph))
			b.Unmarshal(tt.args.graph)
			if got, _ := b.FindShortestPath(tt.args.src, tt.args.dst); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BFSGraph.BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
