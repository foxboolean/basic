package main

import "fmt"

func main() {
	ifUsingNewVariable(10, 200)
}

func ifUsingNewVariable(start int, end int) {
	if distance := end - start; distance > 100 {
		fmt.Printf("distance is :%d", distance)
	} else {
		fmt.Printf("near")
	}
}
