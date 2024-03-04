package main

// TODO: Paquete que proporciona funciones para mostrar información en la consola de entrada y salida.
import "fmt"

func main() {

	// TODO: Print - Acción de mostrar o imprimir un valor o información en la consola.

	// Println imprime un valor con un salto de línea al final.
	fmt.Println("Hola Mundo!")

	// TODO: símbolo de formato (%) - Como se debe formatear el valor que se imprime.

	// Printf imprime con un formato específico.
	fmt.Printf("Hola Mundo con formato\n")

	// Para imprimir un valor de manera genérica, se usa %v.
	var nombre = "Mayer"
	fmt.Printf("Mi nombre es: %v\n", nombre)

	// Para imprimir una cadena, se usa %s.
	fmt.Printf("Mi nombre es: %s\n", nombre)

	// Para imprimir enteros, se usa %d.
	var edad = 23
	fmt.Printf("Mi edad es: %d\n", edad)

	// Para imprimir flotantes, se usa %f.
	var estatura = 1.75
	fmt.Printf("Mi estatura es: %f\n", estatura)

	// Para imprimir valores booleanos, se usa %t.
	var vivo = true
	fmt.Printf("Vivo: %t\n", vivo)

	// Para imprimir el tipo de dato, se usa %T.
	fmt.Printf("Tipo: %T\n", nombre)
}
