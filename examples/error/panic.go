package main

import "fmt"

func main() {
	defer func() {
		if data := recover(); data != nil {
			fmt.Printf("hello panic %v\n", data)
		}
		fmt.Printf("recover \n")
	}()
	panic("Boom")
	fmt.Printf("here is not go to do")
}