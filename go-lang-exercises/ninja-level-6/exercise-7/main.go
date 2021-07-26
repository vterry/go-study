// Assign a func to a variable, then call that func

package main

import "fmt"

func main() {

	odds := func() {
		result := 0
		s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		for _, v := range s {
			if v%2 != 0 {
				result += v
			}
		}

		fmt.Println(result)
	}

	odds()
}
