package main

import (
	"fmt"
)

func main() {
	x := [5]int{5, 10, 15, 20, 25}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)
}
