package main

import "fmt"

func main() {
	sampleVarFunc(2, 2)
}

//variadic function test
func sampleVarFunc(x ...int) {
	sum := 0

	for _, v := range x {
		sum += v
	}

	fmt.Printf(`%T`, x)
}
