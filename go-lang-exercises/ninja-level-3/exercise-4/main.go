/*
Create a for loop using this syntax
	for { }
Have it print out the years you have been alive.

*/

package main

import "fmt"

func main() {
	by := 1992

	for {
		fmt.Println(by)
		if by <= 2021 {
			by++
			continue
		} else {
			break
		}
	}
}
