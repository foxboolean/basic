package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := make([]int, 2, 4)
	fmt.Printf("%d\n", unsafe.Sizeof(s))

	fmt.Printf("%T", &s)
}