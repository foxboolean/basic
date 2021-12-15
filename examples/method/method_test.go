package method

import (
	"fmt"
	"testing"
)

func TestFollow(t *testing.T) {
	Follow()
}

func TestDeferMethod(t *testing.T) {
	DeferMethod()
}

func TestNextInt(t *testing.T) {
	b:=[]byte("123acaws12")
	x, i:=NextInt(b, 1)
	fmt.Printf("val %d, next index is %d\n", x, i)
}