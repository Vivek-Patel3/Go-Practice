package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	defer wg.Wait() // wait must run before cancel otherwise cancel will stop the go-routines and there will be no use of wait then, therefore put it after defer cancel() OR don't use defer for wg.Wait()

	mockInput := func(data interface{}, stream chan interface{}) {
		for {
			select {
			case stream <- data:
			case <-ctx.Done():
				return
			}
		}
	}

	appleStream := make(chan interface{})
	go mockInput("apple", appleStream)

	orangeStream := make(chan interface{})
	go mockInput("orange", orangeStream)

	peachStream := make(chan interface{})
	go mockInput("peach", peachStream)

	// now we have the input. Need to direct this to child go-routines which will perform the processing part

	// 1st child go-routine
	wg.Add(1)
	go func(stream chan interface{}) {
		defer wg.Done()

		var childWg sync.WaitGroup
		defer childWg.Wait()

		childCtx, childCancel := context.WithTimeout(ctx, time.Second * 3)
		defer childCancel()

		// now this will spin up 3 new goroutines to pass on the tasks to them
		for i:=0;i<3;i++ {
			childWg.Add(1)
			go func() {
				defer childWg.Done()
				for {
					select {
					case <-childCtx.Done():
						return
					case val, ok := <-stream:
						if !ok {
							fmt.Println("Input stream closed. No more tasks")
							return
						}
						time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) 
						fmt.Println(val)
					}
			}
			}()
		}

	}(appleStream)

	// now calling the next 2 go-routines
	func2 := genericFunc
	func3 := genericFunc
	
	wg.Add(1)
	go func2(ctx, &wg, orangeStream)
	
	wg.Add(1)
	go func3(ctx, &wg, peachStream)
}

func genericFunc(ctx context.Context, wg *sync.WaitGroup, stream chan interface{}) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case val, ok := <-stream:
			if !ok {
				fmt.Println("Input stream closed. No more tasks")
				return
			}

			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) 
			fmt.Println(val);
		}
	}
}
