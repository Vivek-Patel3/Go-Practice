package main

import (
	"fmt"
	"math/rand"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Vivek-Patel3/coding-task/internal/event"
)

var m sync.Mutex
var storageDisk map[int]string

func process(event event.Event) {
	// converting the payload to uppper case
	event.Payload = strings.ToUpper(event.Payload)

	fmt.Println("Processed payload:", event.Payload)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	// write to the in-memory map
	store(event)
}

func store(event event.Event) {
	m.Lock()
	defer m.Unlock()

	storageDisk[event.Id] = event.Payload
}

func main() {
	eventsChan := make(chan event.Event, 10)
	processChan := make(chan event.Event, 10)
	storageDisk = make(map[int]string)

	var wg sync.WaitGroup

	// producer - produces infinite events
	// one producer -> events for lifetime
	go func() {
		for i:=0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			eventsChan <- event.Event {
				Id: i,
				Payload: "hello from " + strconv.Itoa(i),
			}
		}
	}()

	// multiple consumers as goroutines consume these events
	// fan out to consumers
	for i:=0;i<20;i++ {
		go func() {
			for event := range eventsChan {
				// send to processor-channel
				processChan <- event
				fmt.Println("event consumed with payload:", event.Payload)
			}
		}()
	}

	// fan out to the processors
	for i:=0;i<3;i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for event := range processChan {
				process(event)
			}
		}()
	}

	// handling the graceful shutdown logic
	/* Tasks to be performed in graceful shutdown
	1. stop producing new events
	2. close the eventChan and processChan
	3. Drain the channels -> consumers consume the remaining events
						  -> processors finish processing the remaining events
	*/

	wg.Wait()
}