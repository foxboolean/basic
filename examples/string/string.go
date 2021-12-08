package main

import "unicode/utf8"

func main() {
	printString()
}

func printString() {
	str_en := "Hello"
	str_cn := "你好，世界"
	println(len(str_en))	// 输出 5
	println(utf8.RuneCountInString(str_en))	// 输出 5
	println(len(str_cn))	// 输出 15
	println(utf8.RuneCountInString(str_cn))	// 输出 5

	str := "Hello"
	str += " "
	str += "World"
	//string 只能和 string 拼接
	//str += 1
}


