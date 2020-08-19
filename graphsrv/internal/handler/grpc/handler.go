package grpc

import (
	"context"

	"github.com/labstack/gommon/log"
	"github.com/sKudryashov/graphsrv/internal/model"
	"github.com/sKudryashov/graphsrv/internal/repository"
	"github.com/sKudryashov/graphsrv/internal/service"
	"github.com/sKudryashov/graphsrv/pkg/pb"
)

// NewHandler is a handler constructor
func NewHandler(lgr *log.Logger) *Handler {
	strg := repository.NewStorage(lgr)
	return &Handler{
		lgr:          lgr,
		serviceGraph: service.NewGraphService(lgr, strg),
	}
}

// Handler represents a grpc handler
type Handler struct {
	lgr          *log.Logger
	serviceGraph service.Graph
}

// Put puts object in storage
func (h *Handler) Put(ctx context.Context, req *pb.Graph) (*pb.Status, error) {
	ID, err := h.serviceGraph.SaveGraph(ctx, model.NewGraphModel(req))
	if err != nil {
		return &pb.Status{
			ID:   ID,
			Code: pb.Code(1),
		}, err
	}
	return &pb.Status{
		ID:   ID,
		Code: pb.Code(0),
	}, nil
}

// GetLast returns last added graph with distance, calculated for given src and dst points
func (h *Handler) GetLast(ctx context.Context, req *pb.Search) (*pb.Path, error) {
	result, err := h.serviceGraph.FindMinPathInLast(ctx, model.NewSearch(req))
	if err != nil {
		return nil, err
	}
	path := &pb.Path{
		ID:   result.ID,
		Node: make([]uint32, 0, len(result.Nodes)),
	}
	for _, node := range result.Nodes {
		path.Node = append(path.Node, uint32(node))
	}
	return path, nil
}

// Delete removes given graph from the storage
func (h *Handler) Delete(ctx context.Context, gr *pb.Graph) (*pb.Status, error) {
	if err := h.serviceGraph.Delete(ctx, gr.ID); err != nil {
		// for simplicity let's assume http response code are valid in our API as well
		return &pb.Status{
			ID:   gr.ID,
			Code: pb.Code(1),
		}, err
	}
	return &pb.Status{
		ID:   gr.ID,
		Code: pb.Code(0),
	}, nil
}
