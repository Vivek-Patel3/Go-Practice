package topic

import (
	"github.com/Vivek-Patel3/Producer-Consumer/internal/partition"
	"github.com/Vivek-Patel3/Producer-Consumer/internal/processor"
	"github.com/Vivek-Patel3/Producer-Consumer/internal/model"
)

type Topic struct {
	partitions []*partition.Partition
}

func NewTopic(n int, bufferSize int, proc *processor.Processor) *Topic {
	partitions := make([]*partition.Partition, n)

	for i := 0; i < n; i++ {
		partitions[i] = partition.NewPartition(i, bufferSize, proc)
	}

	return &Topic{
		partitions: partitions,
	}
}

func (t *Topic) Send(event model.Event) {
	p_index := int(event.Id) % len(t.partitions)
	t.partitions[p_index].Publish(event) // add event to the channel
}

func (t *Topic) Close() {
	for _, p := range t.partitions {
		p.Close()
	}
}
