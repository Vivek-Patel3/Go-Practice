package main

import (
	"fmt"
	"sync"
)

var ct int
var lock sync.Mutex

var increment func() = func() {
	lock.Lock()
	defer lock.Unlock() // using defer will ensure call happens even in case of panicing
	ct++;
	fmt.Printf("Incrementing counter: %v\n", ct)
}

var decrement func() = func() {
	lock.Lock()
	defer lock.Unlock()
	ct--;
	fmt.Printf("Decrementing counter: %v\n", ct)
}

func callMutex() {
	var wg sync.WaitGroup

	for i:=0;i<5;i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	for i:=0;i<5;i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement()
		}()
	}

	wg.Wait()
}