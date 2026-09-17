// ./go/bin/go run ./demo-go/2-variable.go

package main

import "fmt"

func main() {
	var a = "First variable" // Type inféré

	x := 2 // Type inféré + auto-déclaration

	x++ // Incrément de x

	fmt.Println(a)
	fmt.Println(x)
}