package main

import "fmt"

func main() {
	//* Los arrays en Go tienen una longitud fija, que debe ser conocida en el momento de la declaración.
	/*
		Los elementos del array se acceden mediante índices, donde el primer elemento está en el índice 0
		y el último está en `len(arr)-1`.
	*/

	// Array sin asignación de valores
	// Los elementos del array se inicializan con el valor cero del tipo de dato correspondiente.
	var array5 [2]bool
	fmt.Println("Array sin asignación de valores:", array5)

	// Declaración de un array de enteros con una longitud de 5.
	var array [5]int
	// Asignación de valores a los elementos del array.
	array[0] = 10
	array[1] = 20
	array[2] = 30
	array[3] = 40
	array[4] = 50
	fmt.Println("Array:", array)

	// Ejemplo de uso de la función `len()` para obtener la longitud del array
	fmt.Println("Longitud de array:", len(array))

	// Declaración de un array con valores literales
	// Los valores se proporcionan directamente entre llaves.
	array2 := [3]string{"Mayer", "Andres", "Chaves"}
	fmt.Println("Array literal:", array2)

	// Declaración de un array multidimensional
	// Se definen arrays dentro de arrays para crear una estructura de datos bidimensional.
	array3 := [3][2]int{{10, 20}, {30, 40}, {50, 60}}
	fmt.Println("Array multidimensional:", array3)

	// Array con longitud implícita
	// El tamaño del array se infiere automáticamente basado en la cantidad de elementos proporcionados.
	array4 := [...]int{10, 20, 30}
	fmt.Println("Array con longitud implícita:", array4)

	// Acceso a un elemento del array
	// Los índices comienzan en 0, por lo que array2[1] accede al segundo elemento.
	fmt.Println("Segundo elemento de array2:", array2[1])

	// Modificación de un elemento del array
	// El primer elemento del array se cambia de "Perro" a "Conejo".
	modificar := [3]string{"Perro", "Gato", "Pez"}
	modificar[0] = "Conejo"
	fmt.Println("Array modificado:", modificar)

	// Los arrays en Go se pasan por valor
	// Asignar un array a otra variable crea una copia de todos los elementos del array original.
	var copyArray = array
	fmt.Println("Copia del array:", copyArray)

	// Recorriendo un array con un `for clásico`
	for i := 0; i < len(array); i++ {
		fmt.Println("Elemento:", array[i])
	}

	// Recorriendo un array con for `range`
	for i, v := range array {
		fmt.Println("I:", i, "V:", v)
	}

}
