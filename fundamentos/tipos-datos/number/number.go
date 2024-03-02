package main

import "fmt"

func main() {

	// TODO: Si solo ponermos uint o int, el valor dependera del sistema operativo, si es 32 bits o 64 bits.

	// Números enteros positivos
	var number uint8 = 10

	fmt.Printf("Tipo: %T - Valor: %v", number, number)

	// Números enteros negativos y positivos
	var number2 int8 = -10

	fmt.Printf("\nTipo: %T - Valor: %v", number2, number2)

	// Números decimales
	var number3 float32 = 10.5

	fmt.Printf("\nTipo: %T - Valor: %v", number3, number3)

	// Casting de números
	var a int16 = 5000
	// TODO: El valor de b no cambia al hacer el casting
	var b uint8 = 10

	// Convierte un valor de un tipo a otro - b es de 8 bits y se castea a 16 bits
	c := int16(b) + a

	fmt.Printf("\nCasting: %T - Valor: %v", c, c)
}
