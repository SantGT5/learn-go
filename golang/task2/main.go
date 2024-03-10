package main

import (
	"fmt"
	"strings"
)

var input = [][]int{
	{0, 1, 2, 3, 45},
	{10, 11, 12, 13, 14},
	{20, 21, 22, 23, 24},
	{30, 31, 32, 33, 34},
}

// output --> "0,1,2,3,45\n10,11,12,13,14\n20,21,22,23,24\n30,31,32,33,34"

func main() {
	result := make([]string, len(input))

	for _, x := range input {
		k := make([]string, len(input))

		for _, y := range x {
			k = append(k, fmt.Sprint(y))
		}

		result = append(result, strings.Join(k, ","))
	}

	fmt.Println(strings.Join(result, "\n"))

	var myMap2 = map[string]uint8{"Adam": 28, "Sarah": 45}

	for i, value := range myMap2 {
		fmt.Printf("my name is: %v\nage: %v\n", i, value)
	}
}
