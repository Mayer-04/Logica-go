package main

import (
	"fmt"
	"slices"
)

func main() {

	// Declaración de un slice - similar a un arreglo pero la longitud no está definida
	// Si declaras un slice pero no le asignas valores explícitos, el slice sera "nil"
	var slice []int

	fmt.Println(slice)

	// Declaración literal - Declara e inicializa los valores en la misma línea
	var slice2 = []int{1, 2, 3, 4}

	fmt.Println(slice2)

	// Declaración literal con variable corta
	chars := []string{0: "a", 2: "c", 1: "b"}

	fmt.Println(chars)

	// Creando un slice a partir de un arreglo
	array := [5]int{1, 2, 3, 4, 5}

	//* slicing operator: Crear nuevos slices basados en otros arrays o slices.
	//TODO: slice comienza en 0, array[:n] es igual a array[0:n]
	slice3 := array[1:3]

	fmt.Println("Slice 3:", slice3)

	// Slices que apunta a los elementos del array desde la posición 1 hasta el final
	slice4 := array[1:]

	fmt.Println("slice4:", slice4)

	//TODO: Funciones integradas para trabajar con slices

	// Tamaño o longitud de un slice - len()
	// Número de elementos que contiene el slice
	fmt.Println("longitud:", len(slice4))

	// Capacidad de un slice - cap()
	// Número de elementos desde el comienzo del slice hasta el final del array subyacente.
	fmt.Println("capacidad:", cap(slice4))

	// Añade uno o más elementos al slice - append()
	//TODO: append modifica el array al que apunta el slice
	slice5 := append(slice4, 6)
	fmt.Println(slice5)

	clear(slice4)
	fmt.Println("clear:", slice4)

	// Desempaquetando un slice o spread operator
	//* Permite que los elementos se pasen individualmente como argumentos utilizando los 3 puntos
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	c := append(a, b...)

	fmt.Println(c)

	// Creando un slice con make
	//TODO: El primer parametro es el tipo de dato, el segundo es la longitud y el tercero la capacidad
	make1 := make([]byte, 5)
	fmt.Println(make1)

	// Go en su versión 1.22 introduce la función `Concat()`
	// Devuelve un nuevo slice que concatena los dos slices proporcionados.
	s1 := []string{"Mayer", "Andres"}
	s2 := []string{"Luis", "Maria"}
	concat := slices.Concat(s1, s2)
	fmt.Println(concat) // [Mayer Andres Luis Maria]

	// Go en su versión 1.23 introduce la función `Repeat()`
	// Devuelve un nuevo slice que repite el segmento proporcionado el número de veces indicado.
	numbers := []int{0, 1, 2, 3}
	repeat := slices.Repeat(numbers, 2)
	fmt.Println(repeat) // [0 1 2 3 0 1 2 3]
}
