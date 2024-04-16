package main

import "fmt"

//TODO: Struct - Una definición de una clase en Go
// Suelen declararse a nivel de paquete y no a nivel de funciones
type Persona struct {
	Nombre string
	Edad   uint8
}

func main() {

	//* Crea una variable "maria" que contiene una estructura (struct) con los valores cero de cada campo
	var maria Persona
	fmt.Println("maria", maria)

	// Creando una variable - Instanciando una persona
	// TODO: Se debe inicializar todos los campos (Los nombres) - RECOMENDADO
	luis := Persona{
		Nombre: "Luis",
		Edad:   25,
	}

	fmt.Printf("Persona1: %+v\n", luis)

	// Definiendo solo un campo
	// TODO: Los campos no son obligatorios - Los demás campos tomaran el valor 0 de su tipo por defecto
	andres := Persona{
		Nombre: "Andres",
	}

	fmt.Println("Persona2:", andres)

	// Asignando un valor inicial en forma ordenada - No necesitamos específicar el nombre de los campos
	mayer := Persona{"Mayer", 23}

	fmt.Println("Persona3:", mayer)

	// Usando el formato campo:valor para inicializarlo sin orden
	lucas := Persona{Edad: 40, Nombre: "Lucas"}

	fmt.Println("Persona4:", lucas)

	// Accediendo a los campos - Usando el operador punto "."
	fmt.Println(luis.Nombre) // Output: Luis
	fmt.Println(luis.Edad)   // Output: 25

	// Struct anónima - Se define y se instancia en el mismo lugar sin tener un nombre definido
	// TODO: Útiles cuando sólo se usan en un contexto específico y no es necesario reutilizarlos
	objeto := struct {
		A bool
		B int
		C bool
	}{
		A: true,
		B: 42,
		C: false,
	}

	fmt.Printf("%+v\n", objeto)

	//* Operador & para generar un puntero al struct "objeto"
	objectoCopy := &objeto
	objectoCopy.B = 24
	fmt.Printf("objeto: %+v\n", objeto)
}
