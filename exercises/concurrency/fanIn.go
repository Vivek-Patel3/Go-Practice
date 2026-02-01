package main

import "fmt"

func fanIn(chan1, chan2 <-chan string) chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-chan1
		}
	}()

	go func() {
		for {
			c <- <-chan2
		}
	}()

	return c
}

func callFanIn() {
	c := fanIn(boring("Instagram"), boring("Facebook"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("You both are too boring. I am quitting")
}