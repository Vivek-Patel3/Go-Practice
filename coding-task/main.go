package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/Vivek-Patel3/coding-task/internal/event"
)

var m sync.Mutex
var storageDisk map[int]string

func process(event event.Event) {
	// converting the payload to uppper case
	event.Payload = strings.ToUpper(event.Payload)

	fmt.Printf("Event: %v processed\n",event.Id)
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

	ctx, cancel := context.WithCancel(context.Background())
	
	var wg sync.WaitGroup
	var consumer_wg sync.WaitGroup

	// producer - produces infinite events
	// one producer -> events for lifetime
	go func(ctx context.Context) {
		for i:=0; ; i++ {
			select {
				case <-ctx.Done():
					return
				case eventsChan <- event.Event {
						Id: i,
						Payload: "hello from " + strconv.Itoa(i),
					}:
			}
		}
	}(ctx)

	// multiple consumers as goroutines consume these events
	// fan out to consumers
	for i:=0;i<20;i++ {
		consumer_wg.Add(1)
		go func() {
			defer consumer_wg.Done()
			for event := range eventsChan {
				// send to processor-channel
				processChan <- event
				fmt.Printf("Event: %v consumed\n", event.Id)
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
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	if sig, ok := <-signalChan; ok {
		// stop the producer
		cancel()

		// close the channels
		close(eventsChan)
		
		// need to wait till all the consumers drain the eventsChan and push those events to processChan
		// therefore need to wait for all consumer's goroutines to finish, so add waitgroup for consumers 
		consumer_wg.Wait()

		// now close the processChan
		close(processChan)
		fmt.Println("Terminating due to", sig)
	}

	wg.Wait()

	for k,v := range storageDisk {
		fmt.Printf("Event %v: %s\n", k, v)
	}
}