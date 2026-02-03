package main

import (
	"fmt"
	"sync"
	"time"
)

type Resource struct {
	CreatedAt time.Time
}

var (
	instance *Resource
	once     sync.Once
)

// Initialization is guaranteed to run exactly once.
func GetResource() *Resource {
	once.Do(func() {
		fmt.Println("Initializing resource...")
		time.Sleep(2 * time.Second)

		instance = &Resource{
			CreatedAt: time.Now(),
		}
	})

	return instance
}

func callOnce() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			res := GetResource()
			fmt.Printf("Goroutine %d got resource created at %v\n", id, res.CreatedAt)
		}(i)
	}

	wg.Wait()
}
