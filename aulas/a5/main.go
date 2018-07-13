package main

import (
	"fmt"
)

func main() {
	tabuada := [10]int{1, 2, 3}
	tab := []int{0, 0, 0, 4, 5, 6}
	user := map[string]string{
		"name": "Teste",
		"nick": "Testudo",
	}

	tabuada[9] = 10

	user["age"] = "22"

	tab = append(tab, 542, 548515, 666)

	fmt.Println(tabuada)
	fmt.Println(tab)
	fmt.Println(len(tab))
	fmt.Println(user)
}
