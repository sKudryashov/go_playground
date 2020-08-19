package grpc

import (
	"context"
	"reflect"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/sKudryashov/graphsrv/pkg/pb"
)

func TestHandler_Put(t *testing.T) {
	type args struct {
		req *pb.Graph
	}
	tests := []struct {
		args    args
		want    *pb.Graph
		name    string
		wantErr bool
	}{
		{
			name: "put one",
			args: args{
				req: &pb.Graph{
					Edges: []*pb.Edge{
						{
							A: 1,
							B: 2,
						},
						{
							A: 2,
							B: 4,
						},
						{
							A: 1,
							B: 3,
						},
						{
							A: 3,
							B: 4,
						},
						{
							A: 4,
							B: 5,
						},
						{
							A: 4,
							B: 6,
						},
						{
							A: 2,
							B: 5,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "put two",
			args: args{
				req: &pb.Graph{
					Edges: []*pb.Edge{
						{
							A: 1,
							B: 2,
						},
						{
							A: 2,
							B: 4,
						},
						{
							A: 1,
							B: 3,
						},
						{
							A: 3,
							B: 4,
						},
						{
							A: 4,
							B: 5,
						},
						{
							A: 4,
							B: 6,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "put three",
			args: args{
				req: &pb.Graph{
					Edges: []*pb.Edge{
						{
							A: 1,
							B: 2,
						},
						{
							A: 2,
							B: 4,
						},
						{
							A: 1,
							B: 3,
						},
						{
							A: 3,
							B: 4,
						},
						{
							A: 4,
							B: 5,
						},
						{
							A: 4,
							B: 6,
						},
						{
							A: 3,
							B: 6,
						},
						{
							A: 6,
							B: 8,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "put four, edge case, full length",
			args: args{
				req: &pb.Graph{
					Edges: []*pb.Edge{
						{
							A: 1,
							B: 2,
						},
						{
							A: 2,
							B: 3,
						},
					},
				},
			},
			wantErr: false,
		},
	}
	lgr := log.New("test")
	lgr.SetLevel(log.ERROR)
	h := NewHandler(lgr)
	var LastAddedID string
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()
			got, err := h.Put(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.Put() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			LastAddedID = got.ID
			if len(got.ID) == 0 || got.GetCode() != 0 {
				t.Errorf("expect len(got.ID) != 0 && got.GetCode() == 0")
			}
		})
	}
	testsSearch := []struct {
		search  *pb.Search
		want    *pb.Path
		name    string
		wantErr bool
	}{
		{
			search: &pb.Search{
				Src: uint32(1),
				Dst: uint32(3),
			},
			want: &pb.Path{
				Node: []uint32{1, 2, 3},
			},
			name:    "search1",
			wantErr: false,
		},
		{
			search: &pb.Search{
				Src: uint32(1),
				Dst: uint32(8),
			},
			want: &pb.Path{
				Node: []uint32{1, 3, 6, 8},
			},
			name:    "search2",
			wantErr: false,
		},
		{
			search: &pb.Search{
				Src: uint32(2),
				Dst: uint32(5),
			},
			want: &pb.Path{
				Node: []uint32{2, 4, 5},
			},
			name:    "search3",
			wantErr: false,
		},
	}
	for _, tt := range testsSearch {
		t.Run(tt.name, func(t *testing.T) {
			path, err := h.GetLast(context.TODO(), tt.search)
			if (err != nil) != tt.wantErr {
				t.Errorf("unexpected error %v", err)
			}
			if !reflect.DeepEqual(path.Node, tt.want.Node) {
				t.Errorf("expect path.Node and tt.want.Node to be equal")
			}
		})
	}
	// LastAddedID
	testsDelete := []struct {
		graph   *pb.Graph
		want    *pb.Status
		name    string
		wantErr bool
	}{
		{
			name: "delete attempt",
			graph: &pb.Graph{
				ID: LastAddedID,
			},
			want: &pb.Status{
				ID:   LastAddedID,
				Code: pb.Code(0),
			},
			wantErr: false,
		},
	}
	for _, tt := range testsDelete {
		t.Run(tt.name, func(t *testing.T) {
			status, err := h.Delete(context.TODO(), tt.graph)
			if (err != nil) != tt.wantErr {
				t.Errorf("unexpected error %v", err)
			}
			if !reflect.DeepEqual(status, tt.want) {
				t.Errorf("expect path.Node and tt.want.Node to be equal")
			}
		})
	}
}
