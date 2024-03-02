package main

import "fmt"

func main() {

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
}
