package main

import "fmt"

func main() {

	// TODO: Slice es un tipo de referencia, por lo que si uno se modifica entonces afectará al resto.

	// Arreglo con 5 elementos donde el tipo es byte
	var array = [5]byte{'a', 'b', 'c', 'd', 'e'}

	// Slices que apunta a los elementos del array del 0 al 3 - sin incluir el elemento 3
	// TODO: slice comienza en 0, array[:n] es igual a array[0:n]
	slice := array[0:3]

	fmt.Println(slice)

	// Slices que apunta a los elementos del array del 1 al final
	slice2 := array[1:]

	fmt.Println(slice2)

	// TODO: funciones integradas para las slice.

	// Tamaño o longitud de un slice - len()
	fmt.Println(len(slice))

	// Capacidad de un slice - cap()
	fmt.Println(cap(slice))

	// Añade uno o más elementos al slice y no devuelve el slice - append
	// TODO: append va a cambiar el array al que apunta el slice
	slice = append(slice, 'f', 'g')

	fmt.Println(slice)

	// Slice literal - Se declaran entre corchetes
	sliceLiteral := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(sliceLiteral)

}
