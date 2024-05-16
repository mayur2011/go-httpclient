package main

import (
	"fmt"
	"sync"
)

func printValue(v string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(v)
}

func main() {

	var wg sync.WaitGroup

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

	wg.Add(len(value))

	for i, v := range value {
		go printValue(fmt.Sprintf("%d %s", i, v), &wg)
	}

	wg.Wait()
	wg.Add(1)
	printValue("Done from main.go gorountine..!", &wg)
}
