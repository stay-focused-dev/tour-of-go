package main

import "fmt"

func fibonacci() func() int {
	v1 := 0
	v2 := 1

	x := func() int {
		ret := v1
		v1, v2 = v2, v1+v2
		return ret
	}

	return x
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
