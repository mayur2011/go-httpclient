package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)
	go updateMessage("XYZ")

	wg.Wait()

	if msg != "XYZ" {
		t.Error("Expected to find XYZ, but its not there.")
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "XYZ"
	printMessage()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, "XYZ") {
		t.Error("Expected to find XYX, but it is not there.")
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut
	if !strings.Contains(output, "Hello, universe!") {
		t.Error("Expected Hello universe, but its not there.")
	}
	if !strings.Contains(output, "Hello, cosmos!") {
		t.Error("Expected Hello cosmos, but its not there.")
	}
	if !strings.Contains(output, "Hello, world!") {
		t.Error("Expected Hello world, but its not there.")
	}

}
