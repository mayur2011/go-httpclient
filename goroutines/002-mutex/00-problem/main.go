package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func main() {
	msg = "Hello, world!"

	wg.Add(2)
	updateMessage("Hello, universe!")
	updateMessage("Hello, cosmos!")
	wg.Wait()

	fmt.Println(msg)

}

/*
Date Race problem
even though we have wg with those two goroutines causing concurrent processing, both are sharing same resource
therefor race condition is happening
go run -race .
*/
