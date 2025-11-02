package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	for _, w := range strings.Fields(s) {
		if v, ok := res[w]; ok {
			res[w] = v + 1
		} else {
			res[w] = 1
		}
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
