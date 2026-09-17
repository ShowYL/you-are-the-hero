// ./go/bin/go run ./demo-go/4-sans-concu.go

package main

import (
	"fmt"
	"sync"
)

func say(word string) {
	for i := 0; i < 10; i++ {
		fmt.Println(word)
	}
}

func main() {
	say("Hello")
	say("world")
}