package main

import "fmt"

func main() {

	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	x := [][]string{a, b}
	fmt.Println(x)

	for i, nx := range x {
		fmt.Println("record: ", i)
		for j, v := range nx {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, v)
		}
	}
}
