package data

import (
	"bytes"
	"fmt"
	"sync"
)

type SyncedBuffer struct {
	lock    sync.Mutex
	buffer  bytes.Buffer
}

func TypeVal() {
	// pointer
	p := new(SyncedBuffer)
	fmt.Printf("val p type is %T\n", p)
	// data
	var v SyncedBuffer
	fmt.Printf("val v type is %T\n", v)
}

func MakeVal() {
	// data
	s := make([]int, 4)
	fmt.Printf("val s is %v, and type is %T \n", s, s)
	i := new([]int)
	fmt.Printf("val s is %v, and type is %T \n", i, i)
	// pointer
	fmt.Printf("pointer is %T \n", &s)
}