package main

import "fmt"

/*
* Operadores aritméticos
Primero se ejecutan las operaciones entre (), *, /, %, +, -
*/

func main() {
	// Suma +
	suma := 10 + 10
	fmt.Println("Suma:", suma)

	// Resta -
	resta := 10 - 5
	fmt.Println("Resta:", resta)

	// Multiplicación *
	multiplicacion := 10 * 5
	fmt.Println("Multiplicación:", multiplicacion)

	// División /
	division := 10 / 5
	fmt.Println("División:", division)

	// Modulo o Resto %
	modulo := 10 % 5
	fmt.Println("Modulo:", modulo)

	// Ejemplo del orden de ejecución de las operaciones.
	// Primero se realiza la operación que esta entre parentesis y luego se multiplica por 2.
	ejemplo := (2 + 3) * 2
	fmt.Println("Ejemplo:", ejemplo)
}
