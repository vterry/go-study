/*
create a func with the identifier foo that
	takes in a variadic parameter of type int
	pass in a value of type []int into your func (unfurl the []int)
	returns the sum of all values of type int passed in
create a func with the identifier bar that
	takes in a parameter of type []int
	returns the sum of all values of type int passed in
*/

package main

import "fmt"

func foo(p ...int) int {
	sum := 0

	for _, v := range p {
		sum += v
	}

	return sum
}

func bar(p []int) int {
	sum := 0

	for _, v := range p {
		sum += v
	}

	return sum
}

func main() {

	p := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Unfurl call: ", foo(p...))
	fmt.Println("Basic call: ", bar(p))
}
