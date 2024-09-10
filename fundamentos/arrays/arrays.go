package main

import "fmt"

/*
* Arrays: Arreglos
Los Arrays en Go tienen un tamaño fijo y almacenan elementos del mismo tipo en ubicaciones de memoria contiguas.

- Los arrays son: arreglos estaticos, tienen un tamaño fijo y se suelen declarar en la pila (stack).
- El compilador de Go conoce la longitud y la capacidad con anticipación de un array.
- `len y cap` son constantes y siempre iguales.
- Los elementos del array se acceden mediante `índices`, donde el primer elemento está en el índice 0
y el último está en `len(arr)-1`.
- La dirección en memoria de un array es la misma dirección del primero elemento del array. (Ej: arr[0])
- En Go cuando pasas un array como argumento de una función, este recibe una copia de ese arreglo,
no un puntero al primer elemento de la matriz.
*/

func main() {
	// Array sin asignación de valores.
	// Los elementos del array se inicializan con el valor cero del tipo de dato correspondiente.
	// En este caso array5 tiene 2 elementos que corresponden al valor cero del tipo de dato `bool`.
	var array5 [2]bool
	fmt.Println("array sin asignación de valores:", array5) // Output: [false false]

	// Declaración de un array de enteros con una longitud de 5.
	var array [5]int
	// Asignación de valores a los elementos del array.
	array[0] = 10
	array[1] = 20
	array[2] = 30
	array[3] = 40
	array[4] = 50
	fmt.Println("array:", array)

	// Ejemplo de uso de la función `len()` para obtener la longitud del array.
	fmt.Println("longitud de array:", len(array))

	// Declaración de un array con valores literales.
	// Los valores se proporcionan directamente entre llaves.
	array2 := [3]string{"Mayer", "Andres", "Chaves"}
	fmt.Println("array literal:", array2)

	// TODO: Cuando declaramos un array literal arr := [3]int{1, 2, 3, 4}, lo que realmente sucede es esto:
	arr := [4]int{}
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4

	// Declaración de un array multidimensional.
	// Se definen arrays dentro de arrays para crear una estructura de datos bidimensional.
	array3 := [3][2]int{{10, 20}, {30, 40}, {50, 60}}
	fmt.Println("array multidimensional:", array3)

	// Array con longitud implícita.
	// El tamaño del array se infiere automáticamente basado en la cantidad de elementos proporcionados.
	array4 := [...]int{10, 20, 30}
	fmt.Println("array con longitud implícita:", array4)

	//* Array con longitud implícita - Inicialización indexada.
	// Inicializar elementos específicos de un arreglo usando sus índices,
	// dejando los elementos no especificados con valores por defecto.
	array6 := [...]int{5: 3}
	fmt.Println("array6:", array6) // Output: [0 0 0 0 3]

	// Acceder a un elemento del array.
	// Los índices comienzan en 0, por lo que array2[1] accede al segundo elemento.
	fmt.Println("segundo elemento de array2:", array2[1])

	// Modificación de un elemento del array.
	// El primer elemento del array se cambia de "Perro" a "Conejo".
	modificar := [3]string{"Perro", "Gato", "Pez"}
	modificar[0] = "Conejo"
	fmt.Println("array modificado:", modificar) // Output: [Conejo Gato Pez]

	// Los arrays en Go se pasan por valor.
	// Asignar un array a otra variable crea una copia de todos los elementos del array original.
	var copyArray = array
	fmt.Println("copia del array:", copyArray)

	// Recorriendo un array con un `for clásico`.
	for i := 0; i < len(array); i++ {
		fmt.Println("elemento:", array[i])
	}

	// Recorriendo un array con for `range`.
	for i, v := range array {
		fmt.Println("i:", i, "v:", v)
	}
}
