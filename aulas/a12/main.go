package main

import "fmt"

func main() {

	u := User {
		"Teste",
		"Babilonia",
		true,
	}

	a := Admin {
		User{
			"Testado",
			"Babilonio",
			true,
		},
		22,
	}

	showUserInfo(u)
	showUserInfo(a)

}

type UsersInterface interface{
	Show() string
}

func showUserInfo(u UsersInterface) {
	fmt.Println(u.Show())
}

type User struct {
	Name string
	Username string
	online bool
}

func (u User) Show() string {
	return fmt.Sprintf("Vai se foder %s", u.Name)
}

type Admin struct {
	User
	age int
}
