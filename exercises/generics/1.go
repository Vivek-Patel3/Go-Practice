package main

import "fmt"

type Box[T any] struct {
  value T
}

func (b *Box[T]) Get() T {
  return b.value
}

func call1() {
	b := Box[int]{10}
	fmt.Println(b.Get())
}
