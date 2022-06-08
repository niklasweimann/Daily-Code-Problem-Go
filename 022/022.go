package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Should all be true:")
	fmt.Println(IsNumber("10"))
	fmt.Println(IsNumber("-10"))
	fmt.Println(IsNumber("10.1"))
	fmt.Println(IsNumber("-10.1"))
	fmt.Println(IsNumber("1e5"))
	fmt.Println("Should all be false:")
	fmt.Println(IsNumber("a"))
	fmt.Println(IsNumber("x 1"))
	fmt.Println(IsNumber("a -2"))
	fmt.Println(IsNumber("-"))
}

func IsNumber(input string) (match bool) {
	match, _ = regexp.MatchString("^([-]?[0-9]e[0-9]*|[-]?[0-9]+[.]?[0-9]*)", input)
	return
}
