package main

import "fmt"

type List[T any] struct {
	head, tail *element[T]
	size int
}

type element[T any] struct {
	next *element[T]
	val T
}

// methods to implement:
// push, all_elements

func (lst *List[T]) Push(v T) {
	if lst.head == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		last := lst.tail
		last.next = &element[T]{val: v}
		lst.tail = last.next
	}
	lst.size++
}

func (lst *List[T]) ToSlice() []T {
	result := make([]T, lst.size)

	for ele := lst.head; ele != nil; ele = ele.next {
		result = append(result, ele.val)
	}

	return result
}

func (lst *List[T]) len() int {
	return lst.size
}

func callLinkedList() {
	lst := List[int]{}

    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    fmt.Println("list:", lst.ToSlice())
}