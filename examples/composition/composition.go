package main

import "fmt"

func main() {
	b := Base{
		Name: "Tom",
	}
	b.name()

	c := Com1{Base: b}
	c.Base.name()
	c.name()
	c.print()
}

type Base struct {
	Name string
}

func (b Base) name() {
	fmt.Printf("this is Base %s \n", b.Name)
}

func (b Base) print()  {
	fmt.Printf("print base info")
	b.name()
}

type Com1 struct {
	Base
}

func (c Com1) name() {
	fmt.Printf("this is Com1 %s\n", c.Name)
	fmt.Printf("this is Com2 %s\n", c.Base.Name)
}