package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println("hello world")

	sayHello()
}

func sayHello() {
	fmt.Println("Print Hello")
}

func TestHello(t *testing.T) {
	t.Log("test hello")
}
