package main

import (
	"basic/core/article3/q/internal"
	"basic/core/article3/q/lib"
	//test "basic/core/article3/q/test/lib"
	. "basic/core/article3/q/test/lib"
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "every", "Who")
	flag.Parse()
}

func main() {
	//fmt.Printf("Hello %s\n", name)
	lib5.Hello(name)
	// 直接父包可以引用
	internal.Hello()

	Hello(name)
}