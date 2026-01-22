package main

import "fmt"

func modifyArray() {
	var a [5]int = [5]int{1,2,3,4,5}

	for i := range a {
		a[i] *= 10
	}

	fmt.Println(a)
	
	// needing only value in idiomatic range
	for _,v := range a {
		fmt.Print(v, " ")
	}
}