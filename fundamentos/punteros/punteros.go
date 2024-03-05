package main

import "fmt"

// Función que toma un puntero y modifica el valor apuntado
func incrementar(puntero *int) {
	*puntero++
}

func main() {

	// Declaración de una variable
	var numero int = 42
	// Declaración de un puntero con el *
	var puntero *int

	// Asignación del puntero
	puntero = &numero

	// Acceso al valor a través del puntero
	fmt.Println("Valor de numero:", *puntero) // Output: Valor de x: 42

	// TODO: Operador desreferenciación *
	//  Acceder al valor almacenado en la dirección de memoria apuntada por un puntero y modificarlo.
	*puntero = 10

	fmt.Println("Nuevo valor de numero:", numero) // Output: Nuevo valor de x: 10

	// TODO: Ejemplo dos: Declarando una variable puntero y asignandole la posición en memoria
	fruit := "🍎"
	var pointer *string = &fruit

	fmt.Println(pointer)

	// TODO: Ejemplo con la función incrementar
	var x int = 10

	// Pasar un puntero a la función incrementar
	incrementar(&x)

	fmt.Printf("Valor de x: %d, Tipo: %T\n", x, x) // Output: Valor de x: 11
}
