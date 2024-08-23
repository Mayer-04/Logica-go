package main

import (
	"fmt"
)

/*
* Genéricos

Es una característica que nos permite definir funciones, tipos y estrucutas que acepten diferentes tipos de datos.
- Introducido en la versión 1.18 de Go.
- La idea de los genéricos es no sacrificar la seguridad que brindan los tipos.
*/

// * Función genérica
// T: Es un tipo de parámetro que permite a la función aceptar cualquier tipo de dato.
// value: Es el valor que se va a imprimir, que puede ser de cualquier tipo.
// No importa el tipo de dato que se le pase, ya que la función utiliza el tipo de parámetro T.
func print[T any](value T) {
	fmt.Println(value)
}

// * Tipo genérico - Estructura genérica
// Person es una estructura genérica que puede contener un campo Age de tipo int o uint.
// La variable de tipo T es un parámetro de tipo que permite a la estructura Person
// aceptar diferentes tipos de datos numéricos enteros para el campo Age.
type Person[T int | uint] struct {
	// Name es un campo de tipo string que almacena el nombre de la persona.
	Name string
	// Age es un campo de tipo T que almacena la edad de la persona.
	Age T
}

func main() {

	// Llamando a la función genérica
	print(10)
	print("Hola")
	print(true)

	// Estructura genérica
	person1 := Person[uint]{Name: "Mayer", Age: 24}
	fmt.Println(person1)
	person2 := Person[int]{Name: "Mayer", Age: 24}
	fmt.Println(person2)
}
