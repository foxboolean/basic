package main

import (
	"errors"
	"flag"
	"fmt"
)

type Name string

func (n *Name) String() string {
	return fmt.Sprint(*n)
}

func (n *Name) Set(s string) error {
	if len(*n) > 0 {
		return errors.New("name has been setting")
	}
	*n = Name(s + " setted")
	return nil
}

func main() {
	var name Name
	flag.Var(&name, "n", "fox")
	flag.Parse()
	fmt.Println("name:", name)
}
