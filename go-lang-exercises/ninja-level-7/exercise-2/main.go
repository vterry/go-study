package main

import "fmt"

type Person struct {
	firstName string
}

func changeMe(p *Person) {
	p.firstName = p.firstName + " Changed"
}

func main() {
	p1 := Person{"James Bond"}
	fmt.Println("before changeMe(): ", p1.firstName)
	changeMe(&p1)
	fmt.Println("after changeMe(): ", p1.firstName)
}
