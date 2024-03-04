package main

import "fmt"

func main() {

	// Arreglo
	var array [5]int
	array[0] = 10
	array[1] = 20
	array[2] = 30
	array[3] = 40
	array[4] = 50
	fmt.Println(array)

	// Arreglos literales - Se declaran los valores entre corchetes
	array2 := [3]string{"Mayer", "Andres", "Chaves"}

	fmt.Println(array2)

	// Arreglos multidimensionales
	array3 := [3][2]int{{10, 20}, {30, 40}, {50, 60}}

	fmt.Println(array3)

	// Arreglo de longitud implícita
	// TODO: Infiere el tamaño basado en los valores que se le asignen
	array4 := [...]int{10, 20, 30}

	fmt.Println(array4)

	// Arreglos sin asignación de valores
	// TODO: Los elementos del arreglo tendrán el valor cero de su tipo de dato correspondientes
	array5 := [2]bool{}

	fmt.Println(array5)

}
