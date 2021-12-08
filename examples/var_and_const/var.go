package main

// Global 全局可以访问
var Global = "全局变量"

// local 包变量
var local = "包变量"

var (
	First string = "abc"
	second int = 1
)

func main() {
	// Go 可以自己推断所以 int 是灰色
	var a int = 1
	println(a)
	// 省略了类型
	var b = 2
	println(b)
	// uint 不可以省略，默认是 int 类型
	var c uint = 14
	println(c)

	// 编译不通过，go 是强类型语言，不会转换
	//println(a == c)
	// 只声明不赋值
	var d int
	println(d)
	println(Global)
	println(local)
	println(First)
	println(second)
}