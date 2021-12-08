package main

func main() {
	ans := fibonacci(5)
	println(ans)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
