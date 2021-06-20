//Create a program that shows the “if statement” in action.

package main

import "fmt"

func main() {
	bool := 10 > 100

	if bool {
		fmt.Println("I'm here!")
	} else {
		fmt.Printf("I'm here!")
	}
}
