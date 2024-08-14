package main

import (
	"fmt"
	"slices"
)

func main() {

	// Declaración de un slice
	// Los slices son estructuras más flexibles que los arrays, con longitud dinámica.
	// Un slice sin valores iniciales será "nil".
	var slice []int
	fmt.Println("Slice sin valores:", slice)

	// Declaración literal de un slice
	// Aquí se declaran e inicializan los valores del slice en una sola línea.
	var slice2 = []int{1, 2, 3, 4}
	fmt.Println("Slice literal:", slice2)

	// Declaración literal con índices explícitos
	// Los índices pueden ser especificados explícitamente para definir elementos en posiciones específicas.
	chars := []string{0: "a", 2: "c", 1: "b"}
	fmt.Println("Slice con índices explícitos:", chars)

	// Crear un slice a partir de un array
	// Se usa el operador de slicing para crear un nuevo slice a partir de un array existente.
	array := [5]int{1, 2, 3, 4, 5}
	slice3 := array[1:3]
	fmt.Println("Slice 3 (array[1:3]):", slice3)

	// Crear un slice desde una posición hasta el final del array
	// El slice incluye todos los elementos desde el índice especificado hasta el final del array.
	slice4 := array[1:]
	fmt.Println("Slice 4 (array[1:]):", slice4)

	//* Funciones integradas para trabajar con slices

	// Tamaño o longitud de un slice
	// La longitud del slice se obtiene con la función `len()`.
	fmt.Println("Longitud de slice4:", len(slice4))

	// Capacidad de un slice
	// La capacidad del slice, que es el número de elementos desde el inicio del slice hasta el final del array
	// subyacente, se obtiene con `cap()`.
	fmt.Println("Capacidad de slice4:", cap(slice4))

	// Añadir elementos a un slice
	// La función `append()` agrega uno o más elementos al slice y puede modificar el array subyacente.
	slice5 := append(slice4, 6)
	fmt.Println("Slice después de append:", slice5)

	// Limpiar un slice con la función `clear`
	// La función clear elimina todos los elementos del slice y lo deja con valores cero.
	clear(slice4)
	fmt.Println("Slice después de clear:", slice4)

	// Desempaquetar un slice (spread operator)
	// Los tres puntos (...) permiten expandir un slice y pasarlo como argumentos individuales a una función.
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := append(a, b...)
	fmt.Println("Slice después de desempaquetar:", c)

	// Crear un slice con make
	// La función `make()` permite crear un slice con una longitud y capacidad inicial especificadas.
	make1 := make([]byte, 5)
	fmt.Println("Slice creado con make:", make1)

	// Función `Concat` (introducida en Go 1.22)
	// Concatena dos slices en un nuevo slice.
	s1 := []string{"Mayer", "Andres"}
	s2 := []string{"Luis", "Maria"}
	concat := slices.Concat(s1, s2)
	fmt.Println("Slice concatenado:", concat) // [Mayer Andres Luis Maria]

	// Función `Repeat` (introducida en Go 1.23)
	// Repite el segmento del slice el número de veces indicado en un nuevo slice.
	numbers := []int{0, 1, 2, 3}
	repeat := slices.Repeat(numbers, 2)
	fmt.Println("Slice repetido:", repeat) // [0 1 2 3 0 1 2 3]
}
