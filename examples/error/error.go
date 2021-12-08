package main

import (
	"errors"
	"fmt"
)

func main() {
	ErrorsPkg()
}

func ErrorsPkg() {
	// 创建一个错误
	err := errors.New("base error")
	// 包一层
	wrappedErr := fmt.Errorf("first wrapped error %w \n", err)
	// 解除
	if err == errors.Unwrap(wrappedErr) {
		fmt.Printf("unwarapped error \n")
	}
	// Is 逐层解除包装
	if errors.Is(wrappedErr, err) {
		fmt.Printf("wrapped is error\n")
	}
}