package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup

// Faites en sorte que les deux fonctions s'exécutent en parallèle en rajoutant un mot clé à deux endroits stratégiques.

func iter(ID int) {
	for i := 0; i < 1000; i ++ {
		fmt.Println("Goroutine", ID , "itération", i)
	}
	wg.Done()
}

func main () {
	wg.Add(2)
	iter(1)
	iter(2)
	wg.Wait()
}