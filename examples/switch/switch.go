package main

import "fmt"

func main() {
	chooseFruit("apple")
	chooseFruit("aa")
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
