package main

import (
	"fmt"
	"time"
)

func main() {
	//efFor()
}

// Reverse a
func Reverse() {
	a := []int8{1, 2, 3, 4, 5, 6}
	// i, j = i++, j-- 非法，++ 是语句而非表达式
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func efFor() {
	for pos, char := range "日本\x80語" { // \x80 在 UTF-8 编码中是一个非法字符
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}
}

func tryFor() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range s {
		fmt.Printf("index is %d, value is %d \n", i, v)
	}

	for i := range s {
		println(i)
	}

	for _, v := range s {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(time.Second)
}
