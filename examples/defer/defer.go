package main

import "fmt"

func main() {
	// bbb\n aaa
	defer func() {
		fmt.Printf("aaa\n")
	}()

	defer func() {
		fmt.Printf("bbb\n")
	}()
}
