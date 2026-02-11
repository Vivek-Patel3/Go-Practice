package partition

import (
	"fmt"
	"github.com/Vivek-Patel3/Producer-Consumer/internal/model"
	"github.com/Vivek-Patel3/Producer-Consumer/internal/processor"
)

type Partition struct {
	id        int
	events    chan model.Event
	processor *processor.Processor
}

func NewPartition(id int, bufferSize int, proc *processor.Processor) *Partition {
	p := &Partition{
		id:        id,
		events:    make(chan model.Event, bufferSize),
		processor: proc,
	}

	go p.startConsumer()

	return p
}

func (p *Partition) startConsumer() {
	for event := range p.events {
		fmt.Println("Partition", p.id, "processing event", event.Id)
		p.processor.Process(event)
	}
}

func (p *Partition) Publish(event model.Event) {
	p.events <- event
}

func (p *Partition) Close() {
	close(p.events)
}