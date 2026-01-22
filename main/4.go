package main

import "fmt"

func multipleValuesInSwitch() {
	var x int
	fmt.Scan(&x)

	switch(x) {
	case 1,7:
		fmt.Println("Weekend")
	case 2,3,4,5,6:
		fmt.Println("Weekday")
	default:
		fmt.Println("Invalid number")
	}
}