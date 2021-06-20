//Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

package main

import "fmt"

func main() {
	bool := 10 > 100

	if bool {
		fmt.Println("I'm here inside the if")
	} else if !bool {
		fmt.Println("I'm here inside the else if")
	} else {
		fmt.Printf("I'm here inside the else")
	}
}
