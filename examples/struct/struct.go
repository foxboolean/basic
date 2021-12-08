package main

import "fmt"

type ToyDuck struct {
	Name string
	Price float32
}

func (t ToyDuck) Swim() {
	fmt.Printf("duck %s is Swiming \n", t.Name)
}

func main() {
	duck1 := ToyDuck{Name: "Tom"}
	duck1.Name = "Tak"
	duck1.Swim()

	duck2 := &ToyDuck{"Tim",2.1}
	duck2.Name = "Tak"
	duck2.Swim()

	duck3 := new(ToyDuck)
	duck3.Swim()
}