package main

import "fmt"

func main() {
	var active bool = true

	// Imprime el formato de la variable - %T es para el tipo de dato y %v es para el valor
	fmt.Printf("Tipo: %T - Valor: %v", active, active)

	var active2 bool = false

	fmt.Printf("Tipo: %T - Valor: %v", active2, active2)

	// active3: Inferiendo el tipo de dato
	var active3 = true
	fmt.Println(active3)
}
