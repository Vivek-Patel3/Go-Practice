package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

func main() {
	name := flag.String("name", "Motadata", "Name to greet")
	flag.Parse()

	fmt.Println("Hello,", *name)

	// printing the environment info
	fmt.Println("\nEnvironment Info:")
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("Go version:", runtime.Version())

	user := os.Getenv("USER")

	// there is no env variable named env
	if(user == "") {
		user = os.Getenv("USERNAME")
		// USER for linux
		// USERNAME for windows
	}

	fmt.Println("User:", user)
}