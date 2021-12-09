package main

import (
	"fmt"
)

func main() {
	compare:= Compare([]byte("aasd"), []byte("aacd"))
	println(compare)
	//breakTest()
	//chooseFruit("apple")
	//chooseFruit("aa")
}
// Compare 比较两个字节型切片，返回一个整数
// 按字典顺序.
// 如果a == b，结果为0；如果a < b，结果为-1；如果a > b，结果为+1
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}

func breakTest() {
	Out:
	for i := 0; i < 10; i++ {
		fmt.Printf("current loop is %d \n", i)
		Loop:
			for j := 0; j < 5; j++ {
				fmt.Printf("inner loop is %d \n", j)
				if j == 0 {
					fmt.Printf("break Loop \n")
					continue Loop
				}
				if j == 1 {
					fmt.Printf("break Out \n")
					break Out
				}
			}
	}
	fmt.Printf("end\n")
}

func chooseFruit(s string) {
	switch s {
	case "apple":
		fmt.Printf("this is an apple\n")
	case "banana":
		fmt.Printf("banana is my favourite\n")
	case "pear":
		fmt.Printf("pear is best\n")
	default:
		fmt.Printf("do not have it \n")
	}
}
