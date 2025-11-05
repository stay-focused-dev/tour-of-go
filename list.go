package main

import (
	"fmt"
	"strings"
)

type List[T any] struct {
	next *List[T]
	val  T
}

func (x *List[T]) Add(v T) *List[T] {
	n := &List[T]{val: v, next: nil}

	if x == nil {
		return n
	}

	c := x
	for c.next != nil {
		c = c.next
	}
	c.next = n

	return x
}

func (x *List[T]) String() string {
	if x == nil {
		return "[]"
	}

	r := []string{}
	c := x
	for c != nil {
		r = append(r, fmt.Sprintf("%v", c.val))
		c = c.next
	}
	return fmt.Sprintf("[%v]", strings.Join(r, ", "))
}

func main() {
	var x *List[int]
	x = x.Add(1)
	x = x.Add(2)
	x = x.Add(3)
	fmt.Println(x)
}
