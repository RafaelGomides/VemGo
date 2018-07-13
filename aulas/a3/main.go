package main

import (
	"fmt"
)

func main() {
	i := 20

	for i > 0 {
		fmt.Println(i)
		i--
	}

	for i := 0; i <= 20; i++ {
		fmt.Println(i)
	}
}
