package main

import "fmt"

func main() {

	// Declaración de un slice - similar a un arreglo pero la longitud no está definida
	var slice []int

	fmt.Println(slice)

	// Declaración literal - Declara e inicializa los valores en la misma línea
	var slice2 = []int{1, 2, 3, 4}

	fmt.Println(slice2)

	// Declaración de slice corta
	chars := []string{0: "a", 2: "c", 1: "b"}

	fmt.Println(chars)

	// Creando un slice a partir de un arreglo
	array := [5]int{1, 2, 3, 4, 5}

	// Slice de la posición 1 a la 3 del arreglo (3 no incluido)
	// TODO: slice comienza en 0, array[:n] es igual a array[0:n]
	slice3 := array[1:3]

	fmt.Println(slice3)

	// Slices que apunta a los elementos del array desde la posición 1 hasta el final
	slice4 := array[1:]

	fmt.Println(slice4)

	// TODO: Funciones integradas para trabajar con slices

	// Tamaño o longitud de un slice - len()
	fmt.Println(len(slice4))

	// Capacidad de un slice - cap()
	fmt.Println(cap(slice4))

	// Añade uno o más elementos al slice - append()
	// TODO: append modifica el array al que apunta el slice
	slice5 := append(slice4, 6)
	fmt.Println(slice5)

	// Desempaquetando un slice o spread operator - ...
	// TODO: Permite que los elementos se pasen individuamente como argumentos
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	c := append(a, b...)

	fmt.Println(c)

	// Creando un slice con make
	//TODO: El primer parametro es el tipo de dato, el segundo es la longitud y el tercero la capacidad
	make1 := make([]byte, 5)

	fmt.Println(make1)

}
