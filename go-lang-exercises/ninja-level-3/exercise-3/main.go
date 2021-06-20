/*
Create a for loop using this syntax
for condition { }
Have it print out the years you have been alive.

*/

package main

import "fmt"

func main() {
	for birthYear := 1992; birthYear <= 2021; birthYear++ {
		fmt.Println(birthYear)
	}
}
