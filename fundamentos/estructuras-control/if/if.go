package main

import "fmt"

//* En Go no existe el operador ternario.
//* En Go no hay conversión automática a booleano
//* La condición dentro de un if debe ser explícitamente una expresión booleana.
//* No se permite asignación en la condición. Esto causará un error de compilación.

func main() {

	//* Estructura if - La condición no debe estar entre paréntesis
	// La condición debe ser verdadera para que se ejecute el bloque de código
	if true {
		fmt.Println("Hola Mundo!")
	}

	edad := 18
	//* Estructura if-else
	// Si la edad es mayor o igual a 18, mostrara el mensaje "Eres mayor de edad"
	// De lo contrario, mostrara el mensaje "Eres menor de edad"
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

	//* If con declaración de variable corta
	// El ámbito de la variable corta 'mayor' es el bloque de código en el que se encuentra
	if mayor := edad >= 18; mayor {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

	//* Estructura if-else
	// Esta estructura permite evaluar múltiples condiciones
	if edad < 13 {
		fmt.Println("Eres un niño")
	} else if edad < 18 {
		fmt.Println("Eres un adolescente")
	} else {
		fmt.Println("Eres un adulto")
	}

}
