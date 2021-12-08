package pkg

import "fmt"

// Main 函数只是普通的打印日志的方法
func Main()  {
	main()
}

// main 函数无法在外部访问
func main() {
	fmt.Println("this is main method in pkg")
}

