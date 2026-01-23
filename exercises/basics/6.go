package main

import "fmt"

func traverseArray() {
	var a [5]int = [5]int{10,20,30,40,50}

	for i,v := range a {
		fmt.Print(i,v," ")
	}

	fmt.Println()
	
	for i := 0;i<len(a);i++ {
		fmt.Print(i,a[i]," ")
	}
}