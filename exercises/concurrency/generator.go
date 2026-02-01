package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Multiplexing two channels into one combined one
func boring(name string) <-chan string {
	c := make(chan string)

	// below goroutine will run forever and will communicate with the channel returned
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s: %v", name, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond) // time.Duration converts the int returned
		}
	}()

	return c
}

func callGenerator() {
	c := boring("Instagram")

	for i:=0;i<5;i++ {
		fmt.Println(<-c)
	}

	fmt.Println("You are too boring. I am quitting")
}