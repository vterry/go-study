/*
Create a func which returns a func
assign the returned func to a variable
call the returned func
*/

package main

import "fmt"

func sample() func() string {
	return func() string {
		return "Running from a returned function"
	}
}

func main() {
	v := sample()
	fmt.Println(v())
}
