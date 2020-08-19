package service

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-uuid"
	"github.com/labstack/gommon/log"
	"github.com/sKudryashov/graphsrv/internal/model"
	"github.com/sKudryashov/graphsrv/internal/repository"
	"github.com/sKudryashov/graphsrv/pkg/graph"
)

// Graph is a business logic layer which represents minPath search by given graph,
// unpacks graph from the transport layer, handles business logic operations with storage/graph search
type Graph struct {
	lgr  *log.Logger
	strg *repository.Storage
}

//NewGraphService is a constructor for service
func NewGraphService(lgr *log.Logger, strg *repository.Storage) Graph {
	return Graph{
		lgr:  lgr,
		strg: strg,
	}
}

// FindMinPathInLast fetches last added graph from the storage and calculates min path between src and dst
func (g *Graph) FindMinPathInLast(ctx context.Context, gm model.SearchModel) (*model.SearchResult, error) {
	g.lgr.Info("received find min path graph request")
	src, dst := gm.GetSrcDst()
	lastAdded, ID, ok := g.strg.GetLast(ctx)
	if !ok {
		return nil, fmt.Errorf("the storage is empty")
	}
	bfsG := graph.NewBFSGraph(lastAdded.Len())
	bfsG.Unmarshal(lastAdded)
	path, err := bfsG.FindShortestPath(src, dst)
	if err != nil {
		return nil, err
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return &model.SearchResult{
		ID:    ID,
		Nodes: path,
	}, nil
}

// Delete deletes graph from the storage
func (g *Graph) Delete(ctx context.Context, ID string) error {
	g.lgr.Info("received delete graph request")
	return g.strg.Delete(ctx, ID)
}

// SaveGraph saves graph into storage
func (g *Graph) SaveGraph(ctx context.Context, gm model.GraphModel) (string, error) {
	g.lgr.Info("received save graph request")
	uuid, err := uuid.GenerateUUID()
	if err != nil {
		return "", err
	}
	if err := g.strg.Insert(ctx, uuid, gm.GetGraph()); err != nil {
		return "", err
	}
	return uuid, nil
}
