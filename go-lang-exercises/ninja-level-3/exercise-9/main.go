// Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.

package main

import "fmt"

func main() {
	favSport := "Dota"

	switch favSport {
	case "Dota":
		fmt.Println("Go to download DOTA on Steam")
	case "Soccer":
		fmt.Println("Go to street with a soccer ball")
	}
}
