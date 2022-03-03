package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.Parse()
	args := flag.Args()
	if len(args) <= 0 {
		return
	}

	switch args[0] {
	case "go":
		// 构造一个新的命令子集 goCmd 绑定参数，最后解析
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "go", "input your name")
		_ = goCmd.Parse(args[1:])
	case "java":
		javaCmd := flag.NewFlagSet("java", flag.ExitOnError)
		javaCmd.StringVar(&name, "name", "java", "input your name")
		_ = javaCmd.Parse(args[1:])
	}
	fmt.Println("name:", name)
}
