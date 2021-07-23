/*
Take the code from the previous exercise, then store the values of type person in a map
with the key of last name. Access each value in the map. Print out the values, ranging over the slice.
*/

package main

import "fmt"

type person struct {
	firstname               string
	lastname                string
	favoriteIceCreamFlavors []string
}

func main() {
	p1 := person{
		"Jotaro",
		"Kujo",
		[]string{"Vanilla", "Tripple Caramel Chunk", "Cookie Dough"},
	}

	p2 := person{
		"Joseph",
		"Joestar",
		[]string{"NewYork Super Fudge Chunck", "Cherry Garcia", "Chocolate Fudge Brownie"},
	}

	m := map[string]person{
		p1.lastname: p1,
		p2.lastname: p2,
	}

	for k, v := range m {
		fmt.Printf("Key %v have the value: %+v \n", k, v)
	}

}
