package main

import "fmt"

func demostratingCopyInArrays() {
	a1 := [4]int{1,2,3,4}
	a2 := a1
	
	a2[0] = 100

	fmt.Println(a1);
	fmt.Println(a2);
}