package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[3:] // len = 2, cap = 2
	s[0] = 10

	fmt.Println(arr)
	s = append(s, 99)
	s[1] = 100

	fmt.Println(arr)
	fmt.Println(s)
	// after exceeding its capacity, s will no longer point to the same underlying array
}