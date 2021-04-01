package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)
	src := []rune(s)
	for left, right := 0, len(src)-1; left < right; left, right = left+1, right-1 {
		src[left], src[right] = src[right], src[left]
	}
	fmt.Println(string(src))
}