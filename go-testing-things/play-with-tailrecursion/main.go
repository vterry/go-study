package main

import (
	"fmt"
)

func fibTailRec(n int) int {

	var accumulator func(index int, last int, previousToLast int) int

	accumulator = func(index int, last int, previousToLast int) int {
		if index >= n {
			return last
		} else {
			return accumulator(index+1, last+previousToLast, last)
		}
	}

	if n <= 2 {
		return 1
	} else {
		return accumulator(3, 2, 1)
	}
}

func main() {
	fmt.Println(fibTailRec(20))
}
