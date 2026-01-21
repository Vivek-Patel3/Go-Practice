package main

import "fmt"

func initializationTypes() {
	var a1 [4]int = [4]int{1,2,3,4}
	a2 := [4]int{1,2,3,4}
	a3 := [4]int{10}
	a4 := [4]int{
		3:40,
	}
	a5 := [...]int{10,20,30,40,50,60}

	fmt.Println(a1);
	fmt.Println(a2);
	fmt.Println(a3);
	fmt.Println(a4);
	fmt.Println(a5);
}