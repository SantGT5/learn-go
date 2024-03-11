package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var input string = "abc"

// expect --> ["ab", "c_"]

func main() {
	var buildStr strings.Builder

	buildStr.WriteString(input)

	var inputLen int = utf8.RuneCountInString(input)

	if inputLen%2 != 0 {
		buildStr.WriteString("_")
	}

	var catStr = buildStr.String()
	var catStrLen int = utf8.RuneCountInString(catStr)

	var res = make([]string, 0, catStrLen/2)

	for i := 0; i < catStrLen; i += 2 {
		res = append(res, catStr[i:i+2])
	}

	fmt.Println(res)
}
