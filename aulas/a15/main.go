package main

import "fmt"

func main() {
	msg := make(chan string)
	done := make(chan bool)

	go sendMessage(msg)
	go receiveMessage(msg, done)

	<-done
}

func sendMessage(msg chan string) {
	msg <- "JOVENS"
}

func receiveMessage(msg chan string, done chan bool) {
	fmt.Println(<-msg)
	done <- true
}
