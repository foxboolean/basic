package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type me struct{
	name string
	age int
}

func (m me) String() string {
	return fmt.Sprintf("name=%s,age=%d", m.name, m.age)
}

func (m *me) Set(s string) error {
	res := strings.Split(s, ";")
	m.name = res[0]
	age, err := strconv.Atoi(res[1])
	m.age = age
	return err
}

var who me

func init() {
	flag.Var(&who, "who", "enter person info")
	flag.Parse()
}
func main() {
	fmt.Printf("%v\n", who)
}