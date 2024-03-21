package main

import "fmt"

var input []int8 = []int8{1, 2, 3, 4, 5, 0, -1, -5}

// output -> [-1, -2, -3, -4, -5, 0, 1, 5]

func main() {
	var result []int8 = make([]int8, 0, len(input))

	for _, v := range input {
		var invert int8 = -v

		result = append(result, invert)
	}

	fmt.Println(result)
}
