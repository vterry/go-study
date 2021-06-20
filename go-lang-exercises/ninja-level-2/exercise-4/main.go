package main

import "fmt"

func main() {
	value := 42
	fmt.Printf("Decimal value is: %d\n Binary value is: %b\n Hex value is: %#x\n", value, value, value)
	newValue := value << 1 //Shift bit one positon to left
	fmt.Printf("Decimal value is: %d\n Binary value is: %b\n Hex value is: %#x\n", newValue, newValue, newValue)
}
