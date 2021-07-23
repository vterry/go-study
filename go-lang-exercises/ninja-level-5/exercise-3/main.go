package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "purple rain",
		},
		luxury: true,
	}

	fmt.Printf("Truck info are - %+v\n", t1)
	fmt.Printf("Sedan info are - %+v\n", s1)

	fmt.Printf("Truck is four wheel - %+v\n", t1.fourWheel)
	fmt.Printf("Sedan number of doors - %+v\n", s1.doors)

}
