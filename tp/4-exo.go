package main

import (
	"fmt"
	"sync"
	"time"
)

var compteA = 100
var compteB = 100

var wg sync.WaitGroup
var muA sync.Mutex
var muB sync.Mutex

// Faites en sorte que le programme se termine en inversant l'ordre de deux lignes de code.

func virementAversB(montant int) {
	muA.Lock()
	time.Sleep(10 * time.Millisecond)
	muB.Lock()

	compteA -= montant
	compteB += montant

	muB.Unlock()
	muA.Unlock()
	wg.Done()
}

func virementBversA(montant int) {
	muB.Lock()
	time.Sleep(10 * time.Millisecond)
	muA.Lock()

	compteB -= montant
	compteA += montant

	muA.Unlock()
	muB.Unlock()
	wg.Done()
}

func main() {
	wg.Add(2)
	go virementAversB(30)
	go virementBversA(50)
	wg.Wait()

	fmt.Println("Compte A :", compteA)
	fmt.Println("Compte B :", compteB)
}
