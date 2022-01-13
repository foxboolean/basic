package main

import (
	"flag"
	"fmt"
	"os"
)

var v string
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init()  {
	cmdLine.StringVar(&v, "name", "every", "who")
	cmdLine.Parse(os.Args[1:])
}
func main() {
	fmt.Printf("Hello, %s!\n", v)
}
