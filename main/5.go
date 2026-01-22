package main

import "fmt"

func withoutConditionSwitch() {
	token := "valid"
	requestCount := 110
	backendUp := true
	
	switch {
	case token != "valid":
		fmt.Println("Unauthorized")
	case requestCount > 100:
		fmt.Println("Too many requests")
	case !backendUp:
		fmt.Println("Service unavailable")
	default:
		fmt.Println("200 OK")
	}
}