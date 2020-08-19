package repository

import (
	"context"
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/sKudryashov/graphsrv/pkg/graph"
)

// NewKVStorage KVStorage constructor
func NewKVStorage() *KVStorage {
	return &KVStorage{
		data: make(map[string]graph.Graph),
	}
}

// KVStorage is an app abstraction over one or more real storages. In our case over KV and Stack,
// which is polyglot-persistence
type KVStorage struct {
	mu   sync.RWMutex
	data map[string]graph.Graph
}

// Insert inserts graph by given ID
func (ds *KVStorage) Insert(ID string, g graph.Graph) {
	ds.data[ID] = g
}

// Len returns storage length
func (ds *KVStorage) Len() int {
	return len(ds.data)
}

// GetByID returns Graph by given ID, or false if is not found
func (ds *KVStorage) GetByID(ID string) (graph.Graph, bool) {
	if val, ok := ds.data[ID]; ok {
		return val, ok
	}
	return graph.Graph{}, false
}

// Delete remove graph by given id
func (ds *KVStorage) Delete(ID string) {
	delete(ds.data, ID)
}

// Storage is an abstract storage layer of any kind of storage, also provides concurrency safe
// operations over concurrent unsafe KVstorage and Stack
type Storage struct {
	lgr         *log.Logger
	dataStorage *KVStorage
	dataStack   graph.Stack
	mu          sync.RWMutex
}

// NewStorage is a storage constructor
func NewStorage(lgr *log.Logger) *Storage {
	return &Storage{
		dataStack:   graph.NewStack(),
		dataStorage: NewKVStorage(),
		lgr:         lgr,
	}
}

// Insert inserts graph into the storage
func (s *Storage) Insert(ctx context.Context, ID string, g graph.Graph) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	s.mu.Lock()
	s.dataStorage.Insert(ID, g)
	s.dataStack.Push(ID)
	s.lgr.Infof("graph with ID %s was, stack length %d, storage length %d", ID, s.dataStorage.Len(), s.dataStack.Len())
	s.mu.Unlock()
	return nil
}

// GetLast returns last graph added to the storage
func (s *Storage) GetLast(ctx context.Context) (graph.Graph, string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	ID, ok := s.dataStack.Pop()
	if !ok {
		s.lgr.Warn("unable to pop from stack")
		return graph.Graph{}, "", false
	}
	if graph, ok := s.dataStorage.GetByID(ID); ok {
		s.lgr.Infof("graph with ID %s was, stack length %d, storage length %d", ID, s.dataStack.Len(), s.dataStorage.Len())
		return graph, ID, true
	} else {
		s.lgr.Errorf("can't find a record with given ID %s, stack length %d, storage length %d", ID, s.dataStack.Len(), s.dataStorage.Len())
		return graph, ID, false
	}
}

// Delete removes graph from the storage by given ID
func (s *Storage) Delete(ctx context.Context, ID string) error {
	s.mu.Lock()
	s.dataStorage.Delete(ID)
	s.dataStack.Remove(ID)
	s.lgr.Infof("graph with ID %s was, stack length %d, storage length %d", ID, s.dataStorage.Len(), s.dataStack.Len())
	s.mu.Unlock()
	return nil
}
