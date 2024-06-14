package main

import "fmt"

//TODO: Struct - Una definición de una "clase" en otros lenguajes de programación en Go
// Suelen declararse a nivel de paquete y no a nivel de funciones
type Persona struct {
	Nombre string
	Edad   uint8
}

func main() {

	//* Crea una variable "maria" que contiene una estructura (struct) con los valores cero de cada campo
	// Los valores cero para cada tipo de dato son: "" (cadena vacía) para strings, 0 para números y false para boolean
	var maria Persona
	fmt.Println("maria:", maria)

	// Creando una variable e instanciando una persona.
	// TODO: Se recomienda inicializar todos los campos especificando sus nombres (campos)
	luis := Persona{
		Nombre: "Luis",
		Edad:   25,
	}

	fmt.Printf("Persona1: %+v\n", luis)

	// Definiendo solo un campo - Los campos no son obligatorios
	// TODO: Los campos no especificados tomarán el valor cero de su tipo de dato correspondiente.
	andres := Persona{
		Nombre: "Andres",
	}

	fmt.Println("Persona2:", andres)

	//* Asignando un valor inicial en forma ordenada
	// No necesitamos especificar el nombre de los campos, pero deben seguir el orden de definición en el struct
	mayer := Persona{"Mayer", 23}

	fmt.Println("Persona3:", mayer)

	//* Usando el formato campo:valor para inicializarlo sin importar el orden
	lucas := Persona{Edad: 40, Nombre: "Lucas"}

	fmt.Println("Persona4:", lucas)

	//* Accediendo a los campos usando el operador punto "."
	fmt.Println(luis.Nombre) // Output: Luis
	fmt.Println(luis.Edad)   // Output: 25

	//* Struct anónimo: se define y se instancia en el mismo lugar sin tener un nombre definido.
	// Útiles cuando solo se usan en un contexto específico y no es necesario reutilizarlos.
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

	//* Operador & para generar un puntero al struct "objeto".
	// Los cambios realizados a través del puntero afectan al struct original.
	objectoCopy := &objeto
	objectoCopy.B = 24
	fmt.Printf("objeto: %+v\n", objeto)
}
