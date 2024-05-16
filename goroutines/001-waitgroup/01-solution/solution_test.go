package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printValue(t *testing.T) {
	// print something to console
	stdOut := os.Stdout

	r, w, _ := os.Pipe() // r for read, w for write and last one is ignored
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go printValue("X", &wg)
	wg.Wait()

	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "X") {
		t.Errorf("Expected to find X, buts its not there")
	}
}
