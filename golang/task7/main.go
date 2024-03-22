package main

import (
	"fmt"
	"strconv"
)

var input string = "1234"

// output --> 1234

func main() {
	num, _ := strconv.Atoi(input)

	fmt.Println(num)
}
