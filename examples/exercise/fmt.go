package main

import "fmt"

func main() {
	str := printNumWith2(2.331)
	println(str)
}

// 输出两位小数
func printNumWith2(float642 float64) string {
	return fmt.Sprintf("%.2f",float642)
}

func printBytes(data []byte) string {
	return fmt.Sprintf("%#x", data)
}
