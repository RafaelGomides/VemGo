package main

import "fmt"

type User struct {
	Name string
	Age int
	Gender bool
}


func main() {
	u := User{
		Age: 22,
		Name: "Teste",
		Gender: true,
	}
	u.aniversario()
	fmt.Println(u)
}

func (u *User) aniversario() {
	fmt.Println("SE FODEU!")
	u.Age++
}
