package main

import (
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "every", "Who")
	flag.Parse()
}

func main() {
	//fmt.Printf("Hello %s\n", name)
	hello(name)
	//lib5.Hello(name)
	// 无法引用
	//internal.Hello()
}