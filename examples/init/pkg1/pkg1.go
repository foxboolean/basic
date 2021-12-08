package pkg1

import "fmt"

const (
	c1 ="c1"
)

var (
	_ = constFmt()
	v1 = varInit("v1")
)

func constFmt() string{
	fmt.Println(c1)
	return c1
}

func varInit(name string) string {
	fmt.Printf("var is %s \n", name)
	return name
}

func init() {
	println("pkg1 init")
}

