package main

import "fmt"

// TODO: Struct - Una definición de una clase en Go
// Suelen declararse a nivel de paquete y no a nivel de funciones
type Persona struct {
	Nombre string
	Edad   uint8
}

func main() {

	// Creando una variable - Instanciando una persona
	// TODO: Se debe inicializar todos los campos - RECOMENDADO
	luis := Persona{
		Nombre: "Luis",
		Edad:   25,
	}

	fmt.Println("Persona1:", luis)

	// Definiendo solo un campo
	// TODO: Los campos no son obligatorios - Los demás campos tomaran el valor 0 de su tipo por defecto
	andres := Persona{
		Nombre: "Andres",
	}

	fmt.Println("Persona2:", andres)

	// Asignando un valor inicial en forma ordenada
	mayer := Persona{"Mayer", 23}

	fmt.Println("Persona3:", mayer)

	// Usando el formato campo:valor para inicializarlo sin orden
	lucas := Persona{Edad: 40, Nombre: "Lucas"}

	fmt.Println("Persona4:", lucas)

	// Accediendo a los campos - Usando el operador punto
	fmt.Println(luis.Nombre) // Output: Luis
	fmt.Println(luis.Edad)   // Output: 25
}
