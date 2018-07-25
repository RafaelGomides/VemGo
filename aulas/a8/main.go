package main

import "fmt"

var i = 10

func main() {
	fmt.Println(soma(7,8,9,100,9000))
	fakeLoop()
}

func soma (numbers ... int) int {
	var total int

	for _, v := range numbers {
		total += v
	}
	return total
}

func fakeLoop(){
	fmt.Println(i)
	i--
	if i > 0 {
		fakeLoop()
	}
}
