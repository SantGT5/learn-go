package main

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"
)

const input string = "1 2 3 4 5"

// expected --> "5 1"

func main() {
	var x []int

	if utf8.RuneCountInString(input) == 0 {
		var err error = errors.New("empty string")
		log.Fatal(err.Error())
	}

	for _, v := range strings.Split(input, " ") {
		num, err := strconv.Atoi(v)

		if err != nil {
			log.Fatal(err)
		}

		x = append(x, num)
	}

	var result string = fmt.Sprint(slices.Max(x), slices.Min(x))

	fmt.Println(result)
}
