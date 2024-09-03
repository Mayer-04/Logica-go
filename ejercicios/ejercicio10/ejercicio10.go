package main

import (
	"fmt"
	"sort"
)

/*
* Encontrar el elemento más grande:
Implementa una función que encuentre el elemento más grande en un slice de números enteros.
La función debería devolver el valor del elemento más grande y su índice en el slice.
*/

func main() {
	enteros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 80, 78}

	mayor, indice := enteroMayor(enteros)
	fmt.Println("número mayor:", mayor)
	fmt.Println("Índice:", indice)

	// Ejemplo 2
	mayor2, indice2 := example2(enteros)
	fmt.Println("número mayor example2:", mayor2)
	fmt.Println("Índice example2:", indice2)
}

func enteroMayor(enteros []int) (int, int) {

	mayor := enteros[0]
	index := 0

	for i := 0; i < len(enteros); i++ { // Empezamos desde el segundo elemento
		if enteros[i] > mayor {
			mayor = enteros[i]
			fmt.Println(i)
			index = i
		}
	}

	return mayor, index // Devolver el valor más grande y su índice
}

func example2(enteros []int) (int, int) {

	// Se ordena el slice de enteros en orden ascendente.
	// Esto hace que el número mayor siempre sea el ultimo.
	// A partir de Go 1.22, esta función simplemente llama a `slices.Sort`.
	sort.Ints(enteros) // slices.Sort(enteros)

	numeroMayor := enteros[len(enteros)-1]

	// Ultimo índice de enteros.
	indice := len(enteros) - 1

	return numeroMayor, indice
}
