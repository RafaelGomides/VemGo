package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	done := make(chan bool)

	go hardTask(done)

	wg.Add(1)
	go loader(done)

	wg.Wait()
}

func hardTask(done chan<- bool) {
	fmt.Println("Start Task")
	time.Sleep(time.Second * 1)
	done <- true
}

func loader(done <-chan bool) {
	defer wg.Done()

	i := 0
	load := []rune(`|\-/`)

	for {
		select {
		case <-done:
			fmt.Println("\r Done")
			return
		default:
			fmt.Printf("\r")
			fmt.Printf(string(load[i]))
			i++
			if i == len(load) {
				i = 0
			}
		}
	}
}
