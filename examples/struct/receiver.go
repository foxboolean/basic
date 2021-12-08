package main

import "fmt"

func main() {
	u := User{Name: "zs", Age: 12}
	fmt.Printf("user detail %v \n", u)
	u.ChangeName("ls")
	fmt.Printf("user detail %v \n", u)
	u.ChangeAge(18)
	fmt.Printf("user detail %v \n", u)

	u2 := &User{Name: "zs", Age: 12}
	fmt.Printf("user detail %v \n", u2)
	u2.ChangeName("ls")
	fmt.Printf("user detail %v \n", u2)
	u2.ChangeAge(18)
	fmt.Printf("user detail %v \n", u2)

}

type User struct {
	Name string
	Age int8
}

// ChangeName 结构体接收器，user 值不会被修改
func (u User) ChangeName(newName string) {
	u.Name = newName
}
// ChangeAge 指针接收器，可以修改 user值
func (u *User) ChangeAge(age int8) {
	u.Age = age
}
