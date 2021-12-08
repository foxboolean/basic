package main

import "fmt"

func main() {
	// 无法直接调用 Fish 方法
	fake := FakeFish{}
	fake.FakeSwim()
	// 转换成 Fish 后才可以调用
	Fish(fake).Swim()


	fish := Fish{}
	fish.Swim()

	sfish := StrongFakeFish{}
	sfish.Swim()
	// AFish 是别称
	afish := AFish{}
	afish.Swim()
}

type Fish struct {
}

func (f Fish) Swim() {
	fmt.Printf("我是一只鱼\n")
}

// FakeFish 定义一个新类型
type FakeFish Fish

func (f FakeFish) FakeSwim() {
	fmt.Printf("我是一只华强北鱼\n")
}

// StrongFakeFish 定义一个新类型
type StrongFakeFish Fish

func (f StrongFakeFish) Swim() {
	fmt.Printf("我是加强版山寨鱼\n")
}

// AFish Fish 的别称
type AFish = Fish