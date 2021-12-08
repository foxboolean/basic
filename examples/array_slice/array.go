package main

import "fmt"

func main() {
	// 数组的 len 和 cap 相等
	a1 := [3]int{1,3,4}
	fmt.Printf("a1: %v, len: %d, cap: %d", a1, len(a1), cap(a1))

	var a2 [3]int
	fmt.Printf("a2: %v, len: %d, cap: %d", a2, len(a2), cap(a2))
}

