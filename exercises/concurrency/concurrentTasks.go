package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// purpose of this program is to get responses from the urls passed in the cli-argument


// using wg.Done() in this function assumes that this will always be called as goroutine
func sendRequest(url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	fmt.Printf("[%s] Status: %d\n", url, resp.StatusCode)
	return nil
}

func callConcurrentTasks() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: go run <file>.go <url1> <url2> ... <urln>")
		} 
		
	start := time.Now()

	var wg sync.WaitGroup
	for _,url := range os.Args[1:] {
		wg.Add(1)

		// not using closure
		go func(url string) {
			defer wg.Done()

			if err := sendRequest("https://" + url); err != nil {
				log.Printf("request failed for %s: %v\n", url, err)
			}
		}(url)
	}

	wg.Wait()
	totalTime := time.Since(start)
	fmt.Printf("%v\n", totalTime)
}