package _map

import "fmt"

func MapKey()  {
	m := make(map[int]string, 2)
	m[1] ="ss"
	println(m[2])
	m2 := make(map[bool]int, 2)
	println(m2[false])
}

func DeleteKey() {
	m := make(map[string]string, 4)
	m["name"] = "zhangsan"
	m["age"] = "12"
	delete(m, "name")
	fmt.Printf("%+v\n", m)
}