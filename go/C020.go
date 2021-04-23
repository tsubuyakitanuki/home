package main

import "fmt"

func main() {
	var m int
	var p int
	var q int
	var unsold int

	fmt.Scan(&m, &p, &q)
	unsold = m * (100 - p) * (100 - q)
	fmt.Println( float64(unsold) / 10000)
}