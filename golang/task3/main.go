package main

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

const input string = "1 2 3 4 5"

// expected --> "5 1"

func main() {
	inputSlice := strings.Split(input, " ")

	if len(inputSlice) == 0 {
		var err error = errors.New("empty array")
		log.Fatal(err.Error())
	}

	intSlice := make([]int, 0, len(inputSlice))

	for _, v := range inputSlice {
		num, err := strconv.Atoi(v)

		if err != nil {
			log.Fatal(err.Error())
		}

		intSlice = append(intSlice, num)
	}

	max, min := slices.Max(intSlice), slices.Min(intSlice)

	fmt.Printf("%d %d\n", max, min)
}
