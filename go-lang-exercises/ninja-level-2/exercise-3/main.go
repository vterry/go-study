package main

import "fmt"

const (
	typed   string = "sample typed"
	untyped        = 32
)

func main() {
	fmt.Printf("Constant one type is: %T and value is: %v\n", typed, typed)
	fmt.Printf("Constant two type is: %T and value is: %v\n", untyped, untyped)
}
