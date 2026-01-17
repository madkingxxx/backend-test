package skinport

import (
	"sync"

	"github.com/madkingxxx/backend-test/internal/core/skinport"
)

type Inmem struct {
	cache map[string]skinport.Item
	mu    sync.RWMutex
}

func New() *Inmem {
	return &Inmem{
		cache: make(map[string]skinport.Item),
	}
}
