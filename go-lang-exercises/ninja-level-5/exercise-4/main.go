package main

import "fmt"

func main() {

	a := struct {
		isThatAnom bool
	}{
		isThatAnom: true,
	}

	fmt.Println(a)
}
