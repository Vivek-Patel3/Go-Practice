package main

import "fmt"

func A() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)
		}
	}()
	B()
	fmt.Println("A continues")
}

func B() {
	panic("boom")
}

func call15() {
	A()
	fmt.Println("main continues...")
}

// once the stack of A has been unwound because of panic, it does not resume execution even after getting recovered
// recover function will only execute in the defer function, if the defer function has been called because of panic