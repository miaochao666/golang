package main

import "fmt"

func main() {
	u1 := User{}
	fmt.Println(u1)
	u1.Age = 10
	fmt.Println(u1)
	u2 := new(User)
	fmt.Println(u2)
	u3 := User{Name: 100}
	fmt.Println(u3)

	var u4 *User
	fmt.Println(u4)
	fmt.Println(u4.Name)

}
