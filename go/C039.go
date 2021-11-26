package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		input_line string
		slice      []string
		result     = 0
	)
	fmt.Scan(&input_line)
	slice = strings.Split(input_line, "+")
	for _, value := range slice {
		result += strings.Count(value, "<") * 10
		result += strings.Count(value, "/")
	}
	fmt.Println(result)
}
