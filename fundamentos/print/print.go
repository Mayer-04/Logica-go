package main

import "fmt"

// * Print proporciona funciones para mostrar información en la consola de entrada y salida.

func main() {

	// * Print: Acción de mostrar o imprimir un valor o información en la consola.

	// Println() imprime un valor con un salto de línea al final.
	fmt.Println("Hola Mundo!")

	// Símbolo de formato (%) - Como se debe formatear el valor que se imprime.

	// Printf() imprime con un formato específico.
	fmt.Printf("Hola Mundo con formato\n")

	// Para imprimir un valor de manera genérica, se usa %v.
	nombre := "Mayer"
	fmt.Printf("Mi nombre es: %v\n", nombre)

	// Para imprimir una cadena, se usa %s.
	fmt.Printf("Mi nombre es: %s\n", nombre)

	// Para imprimir cadenas formateadas, se usa %q.
	cadena := "Hola Mundo!"
	fmt.Printf("Cadena: %q\n", cadena) // "Hola Mundo!"

	// Para imprimir enteros, se usa %d.
	edad := 23
	fmt.Printf("Mi edad es: %d\n", edad)

	// Para imprimir flotantes, se usa %f.
	estatura := 1.75
	fmt.Printf("Mi estatura es: %f\n", estatura)

	// Para imprimir valores booleanos, se usa %t.
	vivo := true
	fmt.Printf("Vivo: %t\n", vivo)

	// Para imprimir el tipo de dato, se usa %T.
	fmt.Printf("Tipo: %T\n", nombre)

	// Para imprimir valores hexadecimales, se usa %x.
	hexadecimal := 0xff
	fmt.Printf("Hexadecimal: %x\n", hexadecimal)

	// Para imprimir valores octales, se usa %o.
	octal := 012
	fmt.Printf("Octal: %o\n", octal)

	// Para imprimir valores binarios, se usa %b.
	binario := 0b1010
	fmt.Printf("Binario: %b\n", binario)

}
