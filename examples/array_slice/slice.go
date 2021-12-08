package main

import "fmt"

func main() {
	arr := [2]int{1,2}
	printArr(arr)

	ss :=[]int{1,2,3,4,5}
	fmt.Printf("%v \n",ss)
	// 子切片和原切片共享底层数组
	ss1 := ss[1:3]
	ss1[0]= -1
	// 2,3
	fmt.Printf("%v \n",ss[1:3])
	// 2,3,4,5
	fmt.Printf("%v \n",ss[1:])
	// 1,2,
	fmt.Printf("%v \n",ss[:2])

	var s []int
	if s == nil {
		 println("slice default value is nil")
	}
	println(s)

	// 初始化 4 个元素的切片
	s1 := []int{1,2,3,4}
	printSlice(s1)
	s2 := make([]int, 4)
	printSlice(s2)
	// cap 4 len 3
	s3 := make([]int, 3, 4)
	printSlice(s3)
	// 添加元素
	s4 := append(s3, 1)
	printSlice(s4)
	// 下标索引
	println(s4[3])
}

func printSlice(s []int) {
	fmt.Printf("s: %v, len: %d, cap: %d \n", s, len(s), cap(s))
}

func printArr(s [2]int) {
	fmt.Printf("s: %v, len: %d, cap: %d \n", s, len(s), cap(s))
}