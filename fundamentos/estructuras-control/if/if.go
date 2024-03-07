package main

import "fmt"

func main() {

	// TODO: La condición no debe estar entre paréntesis

	// Estructura if - La condición debe ser verdadera para que se ejecute el bloque de código
	if true {
		fmt.Println("Hola Mundo!")
	}

	edad := 18
	// Estructura if-else
	// TODO: Si la edad es mayor o igual a 18, mostrar el mensaje "Eres mayor de edad". De lo contrario, mostrar el mensaje "Eres menor de edad".
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

	// If con declaración de variable corta
	// TODO: El ambito de la variable corta es el bloque de código en el que se encuentra
	if mayor := edad >= 18; mayor {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

}
