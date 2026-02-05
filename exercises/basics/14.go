package main

import "fmt"

func call14() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()

    panic("boom")
}

// recover() ONLY works inside a deferred function