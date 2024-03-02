package main

import "fmt"

func main() {

	var a int = 10

	var b *int = &a

	fmt.Printf("Tipo: %T - Valor: %v - Puntero: %v", b, b, *b)

	*b = 20

	fmt.Printf("\nTipo: %T - Valor: %v - Puntero: %v", b, b, *b)
}
