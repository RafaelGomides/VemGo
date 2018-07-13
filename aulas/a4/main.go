package main

import (
	"fmt"
)

func main() {
	_, result, msg := soma(10, 7)
	fmt.Println(result - 1)
	if msg != "" {
		fmt.Println(msg)
	}
}

func soma(a, b int) (int, int, string) {
	fmt.Println("Soma")
	fmt.Println("Resultado: ", a+b)
	if a+b > 10 {
		return 0, a + b, "Valor maior que 10"
	}
	return 0, a + b, ""
}
