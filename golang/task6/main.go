package main

import (
	"fmt"
)

var input []int = []int{1, 2, 10, 8}

// output --> [8, 10]

func main() {
	a, b := 0, 0

	for _, v := range input {
		if v > b {
			a, b = b, v
		} else if v > a {
			a = v
		}
	}

	fmt.Println([2]int{a, b})
}
