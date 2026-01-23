package main

import "fmt"

func useSwitch() {
	var x int
	fmt.Scan(&x)

	switch x {
	case 1:
		fmt.Println("<=5")
	case 2:
		fmt.Println("<=5")
	case 3:
		fmt.Println("<=5")
	case 4:
		fmt.Println("<=5")
	case 5:
		fmt.Println("<=5")
	default:
		fmt.Println(">5")
	}
}