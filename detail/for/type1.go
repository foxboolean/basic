package main

type T1 int
type t2 struct {
	n int
	m int
}

type I interface {
	M1()
}

type S1 struct {
	T1
	*t2
	I
	a int
	b string
}

type S2 struct {
	T1 T1
	t2 *t2
	I  I
	a  int
	b  string
}