package main

import (
	"fmt"
	"time"
)

func printValue(i int, v string) {
	fmt.Println("index ", i, " & value ", v)
}

func main() {

	value := []string{
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"I",
	}

	for i, v := range value {
		go printValue(i, v)
	}
	time.Sleep(1000)
	fmt.Println("Done from main.go gorountine..!")
}
