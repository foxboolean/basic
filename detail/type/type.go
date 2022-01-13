package main

import "time"

type T struct {
	name string
}

func (t T) M1() {
	println(t.name)
}

func (t *T) M2() {
	println(t.name)
}

func main() {
	s := []*T{{"a"}, {"b"}, {"c"}}
	for _, v := range s {
		go v.M1()
	}
	time.Sleep(time.Second * 3)
	//var t T
	//t.M1()
	//t.M2()
	//T.M1(t)
	//(*T).M2(&t)
}
