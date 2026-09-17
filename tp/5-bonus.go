package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mu sync.Mutex

// Faites en sorte que chaque résultat affiché soit correct en rajoutant deux lignes de code.

func puissance(n, m int, res *int) {
	*res = 1
	time.Sleep(time.Millisecond)
	for i := 0; i < m; i++ {
		*res *= n
	}
	fmt.Println(n, "puissance", m, "vaut", *res)
	wg.Done()
}

func main() {
	var res int

	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			wg.Add(1)
			go puissance(i, j, &res)
		}
	}

	wg.Wait()
}
