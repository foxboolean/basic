package main

import (
	_ "callinit/pkg1"
	_ "callinit/pkg2"
	_ "callinit/pkg3"
)

func init() {
	println("main init")
}
func init() {
	println("main init2")
}

func main() {
	println("main method")
}