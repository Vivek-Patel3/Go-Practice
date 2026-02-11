package main

import (
	"sync"
	"github.com/Vivek-Patel3/Producer-Consumer/internal/topic"
	"github.com/Vivek-Patel3/Producer-Consumer/internal/producer"
	"github.com/Vivek-Patel3/Producer-Consumer/internal/processor"
)

func main() {
	proc := processor.NewProcessor()

	// Create topic with 3 partitions
	t := topic.NewTopic(3, 100, proc)

	var wg sync.WaitGroup

	// 20 producers
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int64) {
			defer wg.Done()
			producer.Producer(id, t)
		}(int64(i))
	}

	wg.Wait()

	// producers are done with producing tasks
	t.Close()
}