package main

import (
	"fmt"
)

func main() {
	print("Olá Mundo") // Forma Errada
	fmt.Println("Olá") // Forma certa
	var (
		name  string
		teste = "Ronaldo"
	)

	const (
		AGE = 15
	)

	teste0 := 1

	teste1, teste2 := 10, "Mito"

	name = "Belial"

	fmt.Println(name)
	fmt.Println(teste)
	fmt.Println(teste0)
	fmt.Println(teste0 + teste1)
	fmt.Println(teste0 - teste1)
	fmt.Println(teste0 * teste1)
	fmt.Println(teste0 / teste1)
	fmt.Println(teste2)
}
