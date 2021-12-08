package main

import (
	"basic/examples/main/pkg"
	"fmt"
	_ "net/http"
	"strings"
)

// 无参数，无返回值
// main 方法必须在 main 包中
func main() {
	// go run main.go 执行
	println("Hello world")
	fmt.Println("引入包")
	flag := strings.Contains("Hello", "ll")
	fmt.Println(flag)
	pkg.Main()
}