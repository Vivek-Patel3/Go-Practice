package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

func slave(done <-chan struct{}, input <-chan string, wg *sync.WaitGroup) {
	for {
		select {
		case <-done:
			fmt.Println("No pending tasks. Exiting...")
			return
		case s, ok := <-input:
			if !ok {
				fmt.Println("Input channel has been closed. Exiting...")
				return
			}
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Printf("Processing the string %s: %s\n", s, strings.ToUpper(s))
		}
	}
}

func cancelTask(done chan struct{}) {
	a := 0
	for a == 0 {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		a = rand.Intn(2)
	}
	fmt.Println("Cancelling")
	close(done) // or <-done
}

func callWorkerPool() {
	done := make(chan struct{})
	input := make(chan string)
	var wg sync.WaitGroup

	// starting the worker go-routines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func () {
			defer wg.Done()
			slave(done, input, &wg)
		}()
	}
	
	// simulating cancellation
	go cancelTask(done)
	
	// now begin giving the inputs
	for i:=0;i<100;i++ {
		select {
		case <-done:
			break
		case input <- "hello from " + strconv.Itoa(i):
		}
	}
	
	close(input) // inputs have been completed
	
	// wait for workers to finish
	wg.Wait()
}
