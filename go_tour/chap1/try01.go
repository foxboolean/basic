package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "Go_Tour", "cmd name")
	flag.StringVar(&name, "n", "GG", "simple name")
	flag.Parse()
	fmt.Println("name:", name)
}
