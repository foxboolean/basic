package main

import "fmt"

func main() {
	s := []int{1, 2, 4, 7}
	// 结果应该是 5, 1, 2, 4, 7
	s = Add(s, 0, 5)

	// 结果应该是5, 9, 1, 2, 4, 7
	s = Add(s, 1, 9)

	// 结果应该是5, 9, 1, 2, 4, 7, 13
	s = Add(s, 6, 13)

	// 结果应该是5, 9, 2, 4, 7, 13
	s = Delete(s, 2)

	// 结果应该是9, 2, 4, 7, 13
	s = Delete(s, 0)

	// 结果应该是9, 2, 4, 7
	s = Delete(s, 4)

}

func Add(s []int, index int, value int) []int {
	var res []int
	res = append(res, s[:index]...)
	res = append(res, value)
	res = append(res, s[index:]...)
	fmt.Printf("%v\n", res)
	return res
}

func Delete(s []int, index int) []int {
	var res []int
	res = append(res, s[:index]...)
	if index < len(s) - 1 {
		res = append(res, s[index+1:]...)
	}
	fmt.Printf("%v\n", res)
	return res
}
