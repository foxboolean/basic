package main

type Node struct {
	// 自引用只能用指针
	left *Node
	right *Node
}