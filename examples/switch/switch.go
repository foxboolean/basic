package main

import "fmt"

func main() {
	breakTest()
	//chooseFruit("apple")
	//chooseFruit("aa")
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
					break Loop
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
