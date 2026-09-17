package main

import "fmt"

// Un canal (chan) permet à deux threads de s'échanger des valeurs :
//   c <- v   envoie v (bloque tant que personne ne lit)
//   v := <-c lit v (bloque tant que personne n'écrit)
// Doc : https://go.dev/tour/concurrency/2

// Faites en sorte que le programme se termine en rajoutant une ligne de code.

var done = make(chan bool)

func ping() {
	for i := 0; i < 5; i++ {
		fmt.Println("Ping", i)
	}
}

func main() {
	go ping()
	<-done
	fmt.Println("Fin")
}
