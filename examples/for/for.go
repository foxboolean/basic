package main

import (
	"fmt"
	"time"
)

func main() {
	s:=[]int{1,2,3,4,5,6,7,8,9}
	for i, v := range s {
		fmt.Printf("index is %d, value is %d \n", i, v)
	}

	for i := range s {
		println(i)
	}

	for _, v := range s {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(time.Second)
}
