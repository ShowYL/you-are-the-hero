package main

import (
	"fmt"
	"sync"
	"time"
)

const n = 5

var fourchettes [n]sync.Mutex
var serveur sync.Mutex
var wg sync.WaitGroup

// Faites en sorte que tous les philosophes mangent en rajoutant deux lignes de code .

func philosophe(id int) {
	gauche, droite := id, (id+1)%n

	for repas := 1; repas <= 3; repas++ {
		fourchettes[gauche].Lock()
		time.Sleep(10 * time.Millisecond)
		fourchettes[droite].Lock()

		fmt.Println("Philosophe", id, ": repas", repas)

		fourchettes[droite].Unlock()
		fourchettes[gauche].Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(n)
	for id := 0; id < n; id++ {
		go philosophe(id)
	}
	wg.Wait()

	fmt.Println("Tout le monde a mangé")
}
