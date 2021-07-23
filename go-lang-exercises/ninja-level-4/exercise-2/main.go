package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 8, 9, 10, 11, 12}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)

	y := []string{"Sample", "String", "Composite", "Literal"}

	for i, v := range y {
		fmt.Println(i, v)
	}

}
