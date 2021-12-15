package slice

import "fmt"

type ByteSlice []byte
func (p *ByteSlice) Append(data []byte) {
	slice := append(*p, data...)
	*p = slice
}

func PrintNil() {
	a := [2]int{1,2}
	a2 := [2]int{1,2}
	fmt.Println(a == a2)
	// 切片不能通过 == 比较
	//s := make([]int8, 4)
	//s2 := make([]int8, 4)
	//fmt.Println(s == s2)

	var s []int
	fmt.Printf("s is %v, cap is %d, len is %d \n", s, cap(s), len(s))
}

// LinesOfText 包含多个字节切片的一个切片。
type LinesOfText [][]byte

func PrintLines() {
	text := LinesOfText{
		[]byte("Now is the time"),
		[]byte("for all good gophers"),
		[]byte("to bring some fun to the party."),
	}
	fmt.Printf("%v \n", text)
}

func ExSlice() {
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