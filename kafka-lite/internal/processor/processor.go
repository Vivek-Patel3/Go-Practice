package processor

import (
	"math/rand"
	"sync"
	"time"

	"github.com/Vivek-Patel3/Producer-Consumer/internal/model"
)

type Processor struct {
	store map[int64]string
	mu    sync.Mutex
}

func NewProcessor() *Processor {
	return &Processor{
		store: make(map[int64]string),
	}
}

func (p *Processor) Process(e model.Event) {
	p.mu.Lock()
	defer p.mu.Unlock()

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // simulate processing 	
	p.store[e.Id] = e.Payload
}

func (p *Processor) GetStore() map[int64]string {
	p.mu.Lock()
	defer p.mu.Unlock()


	// creating another copy to avoid passing the actual in-mem store
	copy := make(map[int64]string)
	for k, v := range p.store {
		copy[k] = v
	}
	return copy
}