package main

import "fmt"

func main() {
	i:= 13
	a := func() {
		print(i)
	}
	a()

	Delay()
	// 注：ReturnClosure 返回的是一个方法，所以还要加上 ()
	fmt.Printf(ReturnClosure("Tom")())
}

func Delay() {
	fns := make([]func(), 4)
	if fns == nil {
		fmt.Printf("fns is nil")
	}
	for i := 0; i < 4; i++ {
		fns = append(fns, func() {
			println("this is ", i)
		})
	}

	for _, fn := range fns {
		if fn == nil {
			continue
		}
		fn()
	}
}

func ReturnClosure(name string) func() string{
	return func() string {
		return name
	}
}