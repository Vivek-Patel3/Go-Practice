package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// concurrency breaks order by default
// restoring logic: index input tasks and fan out to workers, fan-in the results
// restore order during fan-in of results not during processing which is blocking operation (waiting for 0-indexed task)

// fan-out -> fan-in with order restoration

type Task struct {
	task string
	id int
} 

type Result struct {
	val string
	id int
}

func worker(c <- chan Task, r chan<- Result) {
	for t := range c {
		// simulating variable processing time
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		
		r <- Result{strings.ToUpper(t.task), t.id}
	}
}

func callRestoreSeq() {
	tasks := make(chan Task)
	results := make(chan Result)

	// set up the worker go-routines before tasks are assigned
	for i:=0;i<3;i++ {
		go worker(tasks, results)
	}

	// now input the tasks

	go func() {
		for i:=0;i<5;i++ {
			tasks <- Task{task: fmt.Sprintf("Task: %d", i), id: i}
		}

		close(tasks)
	}() // this needs to be done via go-routine otherwise, sending tasks will be a blocking operation and no task will be able to return result because results are extracted later 
		// will cause circular deadlock if go-routine is not used here

	// get the results
	result := make([]string, 5)
	for i:=0;i<5;i++ {
		r := <-results
		result[r.id] = r.val
	}

	close(results)

	for i:=0;i<5;i++ {
		fmt.Println(result[i])
	}
}