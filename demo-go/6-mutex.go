// ./go/bin/go run ./demo-go/6-mutex.go

package main

import (
	"fmt"
	"sync"
)

var final = 0
var finalMutex = 0

var wg sync.WaitGroup
var mu sync.Mutex

func addToFinal(value int) {
	final += value
	
	mu.Lock()
	finalMutex += value
	mu.Unlock()
	
	wg.Done()
}

func main() {

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go addToFinal(1)
	}

	wg.Wait()

	fmt.Println("Valeur finale :", final)
	fmt.Println("Valeur finale avec sémaphore :", finalMutex)
}