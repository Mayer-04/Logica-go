package main

import "fmt"

/*
* Funciones de impresión en Go
Go proporciona varias funciones para mostrar información en la consola:

- `print`: Imprime un valor o información en la consola sin un salto de línea,
es una función nativa del lenguaje, diferente de las funciones del paquete fmt.
- `fmt.Print`: Imprime un valor o información en la consola sin un salto de línea.
- `fmt.Println`: Imprime un valor o información en la consola y agrega un salto de línea al final.
- `fmt.Printf`: Imprime una cadena formateada utilizando especificadores de formato.
- Especificadores de formato (%) - Indican cómo se debe formatear el valor que se imprime.
*/

func main() {

	// La función nativa `print()` imprime un texto sin un salto de línea.
	print("Mi primer programa en Go")

	// `fmt.Println()` imprime un valor con un salto de línea al final.
	fmt.Println("Hola Mundo!")

	// `fmt.Printf()` imprime una cadena con un formato específico.
	fmt.Printf("Hola Mundo con formato\n")

	// Usando `%v` para imprimir el valor de una variable de manera genérica.
	nombre := "Mayer"
	fmt.Printf("Mi nombre es: %v\n", nombre)

	// Usando `%s` para imprimir una cadena.
	fmt.Printf("Mi nombre es: %s\n", nombre)

	// Usando `%q` para imprimir una cadena entre comillas.
	cadena := "Hola Mundo!"
	fmt.Printf("Cadena: %q\n", cadena) // "Hola Mundo!"

	// Usando `%d` para imprimir un entero.
	edad := 23
	fmt.Printf("Mi edad es: %d\n", edad)

	// Usando `%f` para imprimir un número flotante.
	estatura := 1.75
	fmt.Printf("Mi estatura es: %f\n", estatura)

	// Usando `%t` para imprimir un valor booleano.
	vivo := true
	fmt.Printf("Vivo: %t\n", vivo)

	// Usando `%T` para imprimir el tipo de dato de una variable.
	fmt.Printf("Tipo: %T\n", nombre)

	// Usando `%x` para imprimir un valor en formato hexadecimal.
	hexadecimal := 0xff
	fmt.Printf("Hexadecimal: %x\n", hexadecimal)

	// Usando `%o` para imprimir un valor en formato octal.
	octal := 012
	fmt.Printf("Octal: %o\n", octal)

	// Usando `%b` para imprimir un valor en formato binario.
	binario := 0b1010
	fmt.Printf("Binario: %b\n", binario)

}
