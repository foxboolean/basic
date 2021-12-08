package main

import (
	"fmt"
)

func main() {
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