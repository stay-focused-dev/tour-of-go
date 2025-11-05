package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// type Tree struct {
// 	Left *Tree
// 	Value int
// 	Right *Tree
// }

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	walk(t, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 {
			return false
		}

		if v1 != v2 {
			return false
		}

		if ok1 == false {
			return true
		}
	}
}

func main() {
	fmt.Println("Walking tree.New(1):")
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("\nTesting Same:")
	fmt.Println("tree.New(1) == tree.New(1):", Same(tree.New(1), tree.New(1)))
	fmt.Println("tree.New(1) == tree.New(1):", Same(tree.New(1), tree.New(2)))
}
