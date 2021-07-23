package main

import "fmt"

type person struct {
	firstname              string
	lastname               string
	favoriteIceCreamFlavor []string
}

func main() {
	p1 := person{
		"First",
		"Person",
		[]string{"Vanilla", "Tripple Caramel Chunk", "Cookie Dough"},
	}

	p2 := person{
		"Second",
		"Person",
		[]string{"NewYork Super Fudge Chunck", "Cherry Garcia", "Chocolate Fudge Brownie"},
	}

	fmt.Printf("The person %v is type of %T \n", p1.firstname, p1)
	fmt.Printf("%v favorite ice cream flavors are: \n", p1.firstname)
	for _, v := range p1.favoriteIceCreamFlavor {
		fmt.Printf("\t%v\n", v)
	}

	fmt.Printf("The person %v is type of %T \n", p2.firstname, p2)
	fmt.Printf("%v favorite ice cream flavors are: \n", p2.firstname)
	for _, v := range p2.favoriteIceCreamFlavor {
		fmt.Printf("\t%v\n", v)
	}
}
