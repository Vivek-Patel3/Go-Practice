package main

import (
	"fmt"
	"time"
)

// doWork -> receiver of the done channel
func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Time to go home and sleep")
			return
		default:
			fmt.Println("Doing some work")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func receiveAvailable() {
	done := make(chan bool) // just a cancellation mechanism, therefore need not be buffered
	
	go doWork(done)
	time.Sleep(time.Second)

	close(done) // we are not sending any value to the done channel, but closed channels are receive available so it will be selected in the select-case

	/*
		could have used done := make(chan struct{}) because we are not going to pass any data 
	*/

	time.Sleep(time.Second) // let the worker print the exiting message
}