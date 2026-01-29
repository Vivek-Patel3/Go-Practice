package main

import "fmt"

// captured variables are shared
func bad() [3]func() {
	fns := [3]func(){}

	// i := 0
	// new declaration in each loop since Go 1.22
	for i := 0; i < 3; i++ {
		fns[i] = func() { fmt.Print(i, " ") }
	}

	return fns
}

// each func captures a new variable
func good() []func() {
	fns := []func(){}

	for i := 0; i < 3; i++ {
		j := i
		fns = append(fns, func() { fmt.Print(j, " ") })
	}

	return fns
}
