package main

import (
	"fmt"
	"os"
)
type T struct {
	a int
	b float64
	c string
}

func Min(a ...int) int {
	min := int(^uint(0) >> 1)  // 最大的 int
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func main() {
	Min(1,2,3,4)
	s := []int {1,2,3,4}
	Min(s...)


	t := &T{ 7, -2.35, "abc\tdef" }
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
	println()


	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))

	fmt.Print(23,21)

	println()
	str2 := "Hello World"
	fmt.Printf("%#x \n",[]byte(str2))

	name := "Tom"
	age := 18
	// 格式化字符串并返回
	str := fmt.Sprintf("My name is %s and my age is %d \n", name, age)
	println(str)

	// 直接打印，简单程序可以直接用来打印信息
	fmt.Printf("My name is %s and age is %d years old \n", name, age)

	u := replaceHolder()
	u.Age = 19
	fmt.Printf("%v", u)
}

func replaceHolder() user{
	u := user{
		Name: "Jam",
		Age: 18,
	}
	fmt.Printf("v => %v \n", u)
	fmt.Printf("+v => %+v \n", u)
	fmt.Printf("#v => %#v \n", u)
	fmt.Printf("T => %T \n", u)
	return u
}

type user struct {
	Name string
	Age int
}