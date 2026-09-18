package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup

// Faites en sorte que les deux fonctions s'exécutent en parallèle en rajoutant une seule ligne de code.

func fibonacci(n int) {
	fmt.Println("Finbonaci de", n)

	prev := 0
	prevPrime := 1
	
	for i := 0; i < n; i++ {
		current := prev + prevPrime
		prevPrime = prev
		prev = current
	}

	fmt.Println("Resultat de fibonacci(", n, ") est de ", prev)
}

func main () {
	wg.Add(2)
	fmt.Println("Début de l'exécution des deux fonctions")
	go fibonacci(400)
	go fibonacci(1000)
	
	fmt.Println("Attente de la fin d'exécution des deux fonctions ...")
	wg.Wait()
}