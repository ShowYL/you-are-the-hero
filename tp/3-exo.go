package main

import (
	"fmt"
	"sync"
)

var x int

var w sync.WaitGroup
var mu sync.Mutex


// Faites en sorte que x s'incrémente correctement jusqu'à 10000 en rajoutant deux lignes de code.

func increment() {
	for i := 0; i < 10; i++ {
		x++
	}
	w.Done()
}

func main() {
	w.Add(1000)
	
	for i := 0; i < 1000; i++ {
		go increment()
	}
	
	w.Wait()

	fmt.Println("x vaut", x)
}
