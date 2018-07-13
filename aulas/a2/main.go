package main

import (
	"fmt"
)

func main() {
	b := true
	c := 2
	a := 7

	if b {
		fmt.Println("Entrou")
	} else {
		fmt.Println("Else")
	}

	if c > 2 {
		fmt.Println("Toma")
	} else {
		fmt.Println("Toma não")
	}

	switch a {
	case 2:
		fmt.Println("É 2")
	case 7:
		fmt.Println("É 7")
	default:
		fmt.Println("Padrão")
	}
}
