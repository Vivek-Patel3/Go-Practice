package main

import "fmt"

type example struct {
	Name string
}

func (e *example) Getter() {
	fmt.Printf("Hi, I'm %s\n", e.Name)
}

func caller() {
	bot1 := &example{Name: "R2-D2"}
	bot2 := &example{Name: "C-3PO"}

	// --- 1. Creating Method Values ---
	greetAsR2 := bot1.Getter 
	greetAs3PO := bot2.Getter

	// --- 2. Using them like standard functions ---
	// These variables now satisfy the signature: func(string)
	greetAsR2()
	greetAs3PO()

	// --- 3. Demonstrating the Closure Property ---
	// If we use a pointer receiver (*example), changing the 
	// instance's data will be reflected when the closure is called.
	bot1.Name = "Artoo"
	
	fmt.Println("\n--- After Name Change ---")
	greetAsR2() // Still points to bot1, but sees the new name
}

// receiver argument is captured by the method value if it is passed by reference
