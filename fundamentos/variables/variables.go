package main

import "fmt"

func main() {

	// TODO: La declaración de variable corta solo puede usarse dentro del cuerpo de una función.

	// Declarando una variable con su tipo
	var name string = "Mayer"

	fmt.Println("nombre:", name)

	// Declarando variable - Tipo inferido
	var name2 = "Andres"

	fmt.Println("segundo nombre:", name2)

	// Declarando una variable - Asignando un valor
	var lastname string
	lastname = "Chaves"

	fmt.Println("Apellido:", lastname)

	// Definición de múltiples variables.
	var casa1, casa2, casa3 string = "🏠", "🏡", "🏚️"

	fmt.Println(casa1, casa2, casa3)

	// Declaración de variable corta - Infiere automáticamente el tipo
	// TODO: No podemos reasignar un valor con la declaración de variable corta.
	age := 23

	fmt.Println("edad:", age)

	// Agrupar múltiples variables
	var (
		home   = "🏠"
		user   = "🧑🏽"
		animal = "🐈"
	)

	fmt.Println(home, user, animal)

	// Definición de tres variables de tipo number, y sus valores iniciales.
	numero1, numero2, numero3 := 2, 3, 4

	fmt.Println(numero1, numero2, numero3)

}
