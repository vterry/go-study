package main

import (
	"encoding/json"
	"fmt"
)

type Hero struct {
	Name   string
	Powers []string
}

func main() {
	h := Hero{
		Name:   "Super Sample",
		Powers: []string{"Creating bug harder to find.", "Creating bug even harder to find."},
	}

	jsonHero, err := json.Marshal(h)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonHero))

}
