package main

import "fmt"

func main() {

	// En Go, si declaras una variable pero no la usas, obtienes un error y el programa no compila
	// La declaración de variable corta solo puede usarse dentro del cuerpo de una función

	// Declarando e inicializando una variable con su tipo explícito
	var name string = "Mayer"

	fmt.Println("nombre:", name)

	// Declarando una variable con tipo inferido
	// Go infiere que name2 es de tipo string por el valor asignado
	var name2 = "Andres"

	fmt.Println("segundo nombre:", name2)

	// Declarando una variable sin inicializarla y asignando un valor después
	// No es una práctica recomendada porque puede llevar a errores si la variable no se inicializa adecuadamente
	var lastname string
	lastname = "Chaves"

	fmt.Println("Apellido:", lastname)

	// Definición de múltiples variables en una sola línea con el mismo tipo
	var casa1, casa2, casa3 string = "🏠", "🏡", "🏚️"

	fmt.Println(casa1, casa2, casa3)

	// Declaración de variable corta
	// Infiere automáticamente el tipo basado en el valor asignado
	// No se puede volver a declarar una variable existente con := en el mismo ámbito
	age := 23

	fmt.Println("edad:", age)

	// Agrupando múltiples variables en un bloque var
	var (
		home   = "🏠"
		user   = "🧑🏽"
		animal = "🐈"
	)

	fmt.Println(home, user, animal)

	// Declaración de múltiples variables de tipo numérico utilizando la declaración de variable corta
	numero1, numero2, numero3 := 2, 3, 4

	fmt.Println(numero1, numero2, numero3)

}
