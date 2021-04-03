package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	src := []rune(s)
	size := len(src) + 2
	fmt.Println(strings.Repeat("+", size))
	fmt.Println("+"+s+"+")
	fmt.Println(strings.Repeat("+", size))
}