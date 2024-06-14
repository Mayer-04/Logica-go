package main

import "fmt"

// Función que toma un puntero y modifica el valor apuntado
func incrementar(puntero *int) {
	// Operador de desreferenciación * para acceder al valor al que apunta el puntero y modificarlo
	*puntero++
}

//* Puntero: Es una variable que contiene la dirección en memoria de otra variable
// Se declara usando el asterisco * seguido del tipo de la variable a la que apunta - var puntero *int

func main() {

	// Declaración e inicialización de una variable
	var numero int = 42

	// Declaración de un puntero de tipo int.
	var puntero *int

	// Asignación del puntero a la dirección de memoria de la variable numero
	puntero = &numero

	// Acceso al valor a través del puntero utilizando el operador de desreferenciación *
	fmt.Println("Valor de numero:", *puntero) // Output: Valor de numero: 42

	// Modificación del valor almacenado en la dirección de memoria apuntada por el puntero
	*puntero = 10

	// El valor de la variable numero se ha actualizado.
	fmt.Println("Nuevo valor de numero:", numero) // Output: Nuevo valor de numero: 10

	// Ejemplo dos: Declaración de una variable puntero y asignación de la dirección de memoria
	fruit := "🍎"
	var pointer *string = &fruit

	// Imprimir la dirección de memoria almacenada en el puntero
	fmt.Println("Pointer:", pointer) // Output: Pointer: 0xc00008a030
	// Imprimir el valor almacenado en la dirección de memoria apuntada por el puntero
	fmt.Println("Value:", *pointer) // Output: Value: 🍎

	// Ejemplo con la función incrementar
	var x int = 10

	// Pasar un puntero a la función incrementar.
	incrementar(&x)

	// Imprimir el valor de x después de ser incrementado por la función.
	fmt.Printf("Valor de x: %d, Tipo: %T\n", x, x) // Output: Valor de x: 11, Tipo: int
}
