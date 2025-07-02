package main

import "fmt"

func SliceIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type LinkedList[T any] struct {
	head, tail *Node[T]
}

type Node[T any] struct {
	next *Node[T]
	val  T
}

func (list *LinkedList[T]) Push(v T) {
	if list.tail == nil {
		list.head = &Node[T]{val: v}
		list.tail = list.head
	} else {
		list.tail.next = &Node[T]{val: v}
		list.tail = list.tail.next
	}
}

func (list *LinkedList[T]) AllElements() []T {
	var elems []T
	for e := list.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}
	fmt.Println("index of zoo:", SliceIndex(s, "zoo"))

	_ = SliceIndex[[]string, string](s, "zoo")

	list := LinkedList[int]{}
	list.Push(10)
	list.Push(13)
	list.Push(23)
	fmt.Println("list:", list.AllElements())
}
