// ./go/bin/go run ./demo-go/5-concu.go

package main

import (
	"fmt"
	"sync"
)

// Variable globale de conteur de synchronisation des threads
var wg sync.WaitGroup

func say(word string) {
	for i := 0; i < 10; i++ {
		fmt.Println(word)
	}
	
	wg.Done() // Indique que le thread est terminé
}

func main() {
	wg.Add(2) // Ajoute 2 threads à attendre
	
	// Ajout du mot-clé "go" pour executer la fonction en 
	go say("Hello")
	go say("world")

	wg.Wait() // Attend que tous les threads soient terminés. Cette fonction est bloquante
}