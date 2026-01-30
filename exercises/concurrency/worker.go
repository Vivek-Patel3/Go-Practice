package main

import (
	"fmt"
	"time"
)

func Worker(done <-chan struct{}, jobs <-chan int) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				fmt.Println("jobs channel closed, worker exiting")
				return
			}
			process(job)

		case <-done:
			fmt.Println("cancellation received, worker exiting")
			return
		}
	}
}

func process(job int) {
	fmt.Println("processing job:", job)
	time.Sleep(3 * time.Second)
}


// call the following function f1 in main func
func f1() {
	done := make(chan struct{})
	jobs := make(chan int)

	go Worker(done, jobs)

	for i := 1; i <= 5; i++ {
		jobs <- i
	}

	// cancel the worker
	close(done)
	close(jobs)

	time.Sleep(time.Second)
}
