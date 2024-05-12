package main

import "fmt"

/*
Copiar un slice:
Implementa una función que copie un slice de enteros en otro slice nuevo.
Asegúrate de que los dos slices sean independientes y no compartan la misma referencia de memoria.
*/

func main() {

	enteros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := copiaSlice(enteros)

	fmt.Println("enteros:", enteros)
	fmt.Println("result:", result)
	fmt.Printf("longitud: %d, capacidad: %d", len(result), cap(result))

}

func copiaSlice(enteros []int) []int {

	var copia = make([]int, len(enteros))

	copy(copia, enteros)

	return copia
}
