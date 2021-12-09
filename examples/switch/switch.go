package main

import (
	"fmt"
)

func main() {
	// 断言
	//t := true
	var t *string
	fmt.Println(t)
	//var t interface{}
	switchType(t)
	//var i interface{}= 1
	//t1:= i.(int)
	//fmt.Println(t1)
	//
	//t2, ok:= i.(int)
	//if ok {
	//	fmt.Println("this is int ", t2)
	//}

	//compare:= Compare([]byte("aasd"), []byte("aacd"))
	//println(compare)
	//breakTest()
	//chooseFruit("apple")
	//chooseFruit("aa")
}

func switchType(i interface{}) {
	switch t := i.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T 打印任何类型的 t
	case nil:
		fmt.Printf("nil %t\n", t) // t 是 nil
	case bool:
		fmt.Printf("boolean %t\n", t) // t 是 bool 类型
	case int:
		fmt.Printf("integer %d\n", t) // t 是 int 类型
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t 是 *bool 类型
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t 是 *int 类型
	case *string:
		fmt.Printf("pointer to string \n") // t 是 *string 类型
	}
}

// Compare 比较两个字节型切片，返回一个整数
// 按字典顺序.
// 如果a == b，结果为0；如果a < b，结果为-1；如果a > b，结果为+1
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}

func breakTest() {
Out:
	for i := 0; i < 10; i++ {
		fmt.Printf("current loop is %d \n", i)
	Loop:
		for j := 0; j < 5; j++ {
			fmt.Printf("inner loop is %d \n", j)
			if j == 0 {
				fmt.Printf("break Loop \n")
				continue Loop
			}
			if j == 1 {
				fmt.Printf("break Out \n")
				break Out
			}
		}
	}
	fmt.Printf("end\n")
}

func chooseFruit(s string) {
	switch s {
	case "apple":
		fmt.Printf("this is an apple\n")
	case "banana":
		fmt.Printf("banana is my favourite\n")
	case "pear":
		fmt.Printf("pear is best\n")
	default:
		fmt.Printf("do not have it \n")
	}
}
