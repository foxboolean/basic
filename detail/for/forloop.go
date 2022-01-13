package main

import "fmt"

func main() {
	//loopString()
	i := 42
	p := &i // p 是指向 i 的指针
	*p = 21 // 修改 i 所在地址保存的值为 21

	loopMap()
}

func loopMap() {
	m := map[string]int { "Rob" : 67, "Russ" : 39, "John" : 29,}
	for k, v := range m {
		if k == "Rob" {
			m["Tom"] = 23
		}
		fmt.Printf("key:%s, value: %d \n", k, v)
	}
	fmt.Printf("%v\n", m)
}

func loopString() {
	str := "中文测试"
	for i, v := range str {
		fmt.Printf("index :%d, value: %s\n", i, string(v))
	}
	for i := 0; i < len(str); i++ {
		fmt.Printf("index :%d, value: %d\n", i, str[i])
	}
}