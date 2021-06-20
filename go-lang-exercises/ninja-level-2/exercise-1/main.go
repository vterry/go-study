package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := strconv.Atoi(os.Args[1])
	/*
		%d ---> Prints a number in decimal
		%b ---> Prints a number in binary
		% ---
	*/

	fmt.Printf("%d\t%#b\t%#x", input, input, input)
}
