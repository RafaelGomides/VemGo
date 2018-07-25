package main

import "fmt"

type List[]interface{}

func main(){
	var l List

	l = []interface{}{
		10,
		"Olá",
		1.9,
		false,
	}

	l.init()

	l.show()
}

func (l List) show() {
	fmt.Println(l...)
}

func (l *List) init() {
	*l = []interface{}{
		"Olá",
		false,
	}
}
