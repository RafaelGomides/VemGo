package main

import "fmt"

type User struct {
	Name string
	online bool
}

type Admin struct {
	User
	Rank int
}

func main() {

	var adm Admin

	adm.Rank = 1
	adm.Name = "Teste"
	adm.online = true

	adm.login()
	fmt.Println(adm)

	adm.logout()
	fmt.Println(adm)

}

func (u *User) login() {
	u.online = true
}

func (u * User) logout() {
	u.online = false
}
