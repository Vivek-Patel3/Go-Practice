package main

import (
	"fmt"
	"time"
)

func Report(metrics chan<- string, metric string) bool {
	select {
	case metrics <- metric:
		return true
	default:
		return false
	}
}

func main() {
	metrics := make(chan string, 2)

	
	// reading the recorded metrics
	go func() {
		for m := range metrics {
			fmt.Println("recording metric:", m)
			time.Sleep(2*time.Second)
		}
		}()
		
	// asking to record the metric (if the channel permits, given that its size is 2)
	for i := 1; i <= 6; i++ {
		metric := fmt.Sprintf("metric-%d", i)

		if Report(metrics, metric) {
			fmt.Println(metric, "recorded")
		} else {
			fmt.Println(metric, "dropped")
		}
		time.Sleep(time.Second) // insert rate more than the reading rate
	}

	close(metrics)
	time.Sleep(10*time.Second) // for the purpose of letting the program finish
}