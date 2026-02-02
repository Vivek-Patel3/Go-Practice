package main

import (
	"sync"
	"time"
)

var producer func(*sync.WaitGroup, sync.Locker) = func(wg *sync.WaitGroup, m sync.Locker) {
	// using the mutex provided, lock the critical section - either mutex or RW mutex
	defer wg.Done()
	// lock the critical section 5 times
	for i:=0;i<5;i++ {
		m.Lock()
		m.Unlock()
		time.Sleep(time.Second)
	}
}

var consumer func(*sync.WaitGroup, sync.Locker) = func(wg *sync.WaitGroup, m sync.Locker){
	defer wg.Done()

	m.Lock()
	m.Unlock()
}

func callRWMutex() {
	// calling one producer and different number of consumers (readers)
	var wg sync.WaitGroup
	var lock sync.RWMutex
	var m sync.Mutex
	
	wg.Add(1)
	// producer must be given normal mutex
	go producer(&wg, &m)

	
}