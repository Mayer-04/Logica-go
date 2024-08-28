package main

import (
	"fmt"
	"slices"
)

/*
* Slices: Rebanadas
- Los slices son estructuras más flexibles que los arrays, con longitud dinámica.
*/

func main() {

	// Declaración de un slice.
	// Un slice sin valores iniciales será "nil".
	var slice []int
	fmt.Println("slice sin valores:", slice)

	// Crear un slice nulo y vacío explícitamente.
	// Un slice nulo es igual a nil mientras que un slice vacío tiene una longitud de 0.
	sliceNull := []string(nil)
	fmt.Println("slice nulo:", sliceNull)

	// Declaración literal de un slice.
	// Aquí se declaran e inicializan los valores del slice en una sola línea.
	var slice2 = []int{1, 2, 3, 4}
	fmt.Println("slice literal:", slice2)

	// Declaración literal de un slice vacío.
	// Los slices vacíos tienen una longitud de 0 y su capacidad es 0.
	var sliceEmpty = []int{}
	fmt.Println("slice literal vacío:", sliceEmpty)

	// Declaración literal con índices explícitos.
	// Los índices pueden ser especificados explícitamente para definir elementos en posiciones específicas.
	chars := []string{0: "a", 2: "m", 1: "z"}
	fmt.Println("slice con índices explícitos:", chars) // Output: ["a", "z", "m"]

	// Crear un slice a partir de un array.
	// Se usa el operador de slicing (:) para crear un nuevo slice a partir de un array existente.
	array := [5]int{1, 2, 3, 4, 5}
	slice3 := array[1:3]
	fmt.Println("slice 3 (array[1:3]):", slice3) // Output: [2 3]

	// Crear un slice desde una posición hasta el final del array.
	// El slice incluye todos los elementos desde el índice especificado hasta el final del array.
	slice4 := array[1:]
	fmt.Println("slice 4 (array[1:]):", slice4) // Output: [2 3 4 5]

	//* Funciones integradas para trabajar con slices

	// Tamaño o longitud de un slice.
	// La longitud del slice se obtiene con la función `len()`.
	fmt.Println("longitud de slice4:", len(slice4))

	// Capacidad de un slice.
	// La capacidad del slice, que es el número de elementos desde el inicio del slice hasta el final del array
	// subyacente, se obtiene con `cap()`.
	fmt.Println("capacidad de slice4:", cap(slice4))

	// Añadir elementos a un slice.
	// La función `append()` agrega uno o más elementos al slice y modifica el `array subyacente`.
	slice5 := append(slice4, 6)
	fmt.Println("slice después de append:", slice5)

	// Limpiar un slice con la función `clear()`.
	// La función clear elimina todos los elementos del slice y lo deja con valores cero.
	clear(slice4)
	fmt.Println("slice después de clear:", slice4)

	// Desempaquetar un slice (...).
	// Los tres puntos permiten expandir un slice y pasarlo como argumentos individuales a una función.
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := append(a, b...)
	fmt.Println("slice después de desempaquetar:", c)

	// Crear un slice con make.
	// La función `make()` permite crear un slice con una longitud y capacidad inicial especificadas.
	// La capacidad es un argumento opcional.
	// El slice creado se inicializa con valores cero.
	make1 := make([]byte, 5)
	fmt.Println("slice creado con make:", make1)

	// Función `Concat` (introducida en Go 1.22).
	// Concatena dos slices en un nuevo slice.
	s1 := []string{"Mayer", "Andres"}
	s2 := []string{"Joe", "Rose"}
	concat := slices.Concat(s1, s2)
	fmt.Println("slice concatenado:", concat) // Output: [Mayer Andres Joe Rose]

	// Función `Repeat` (introducida en Go 1.23).
	// Repite el segmento del slice el número de veces indicado en un nuevo slice.
	numbers := []int{0, 1, 2, 3}
	repeat := slices.Repeat(numbers, 2)
	fmt.Println("slice repetido:", repeat) // Output: [0 1 2 3 0 1 2 3]
}
