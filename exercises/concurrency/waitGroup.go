package main

import (
	"fmt"
	"sync"
)

func hello(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Hello from %v\n", id)
}

func callWaitGroup() {
	var wg sync.WaitGroup

	n := 5
	
	wg.Add(n)
	for i:=0;i<n;i++ {
		go hello(&wg, i)
	}

	wg.Wait()
}