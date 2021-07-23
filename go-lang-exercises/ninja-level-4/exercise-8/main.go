package main

import "fmt"

func main() {

	m := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	// for k, v := range m {
	// 	fmt.Println("Records of ", k)
	// 	for i := 0; i < len(v); i++ {
	// 		fmt.Printf("\t index position: %v \t value: \t %v \n", i, v[i])
	// 	}
	// }

	for k, v := range m {
		fmt.Println("Records of ", k)
		for i, sv := range v {
			fmt.Printf("\t index position: %v \t value: \t %v \n", i, sv)
		}
	}

}
