package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	var m float64
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Println(math.Abs((n - m)))
}
