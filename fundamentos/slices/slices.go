package main

import "fmt"

func main() {
	// Arreglo
	var array = [5]int{10, 20, 30, 40, 50}

	// Slices
	slice := array[0:3]

	fmt.Println(slice)

	// Tamaño o longitud de un slice - len()
	fmt.Println(len(slice))

	// Capacidad de un slice - cap()
	fmt.Println(cap(slice))
}
