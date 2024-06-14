package main

import "fmt"

func main() {
	//* Los arrays en Go tienen una longitud fija, definida al momento de su declaración.
	// Para acceder a cada elemento del array se usa la notación de subíndice [].
	// Donde 0 es el primer elemento y el último es la longitud del array menos 1, es decir, arr[len(arr)-1].

	// Declaración de un array de enteros con una longitud de 5.
	var array [5]int
	array[0] = 10
	array[1] = 20
	array[2] = 30
	array[3] = 40
	array[4] = 50
	fmt.Println("Array:", array)

	// Arrays literales - Se declaran los valores directamente entre corchetes.
	array2 := [3]string{"Mayer", "Andres", "Chaves"}
	fmt.Println("Array literal:", array2)

	// Arrays multidimensionales - Arrays dentro de arrays.
	array3 := [3][2]int{{10, 20}, {30, 40}, {50, 60}}
	fmt.Println("Array multidimensional:", array3)

	// Array de longitud implícita - El tamaño se infiere automáticamente basado en el número de elementos.
	array4 := [...]int{10, 20, 30}
	fmt.Println("Array de longitud implícita:", array4)

	// Arrays sin asignación de valores - Los elementos del array tendrán el valor cero de su tipo de dato correspondiente.
	var array5 [2]bool
	fmt.Println("Array sin asignación de valores:", array5)

	// Acceder a un elemento del array - Los índices comienzan desde 0.
	fmt.Println("Primer elemento de array2:", array2[0])

	// Los arrays son valores, asignar un array a otra variable copia todos los elementos.
	var copyArray = array
	fmt.Println("Copia del array:", copyArray)

	// Modificar un elemento de un array - Cambiamos el primer elemento del array.
	modificar := [3]string{"Perro", "Gato", "Pez"}
	modificar[0] = "Conejo"
	fmt.Println("Array modificado:", modificar)
}
