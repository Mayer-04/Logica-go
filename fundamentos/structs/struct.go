package main

import "fmt"

// Una Struct es una colección de campos que se pueden definir en un paquete.
// En Go, las estructuras suelen declararse a nivel de paquete y no a nivel de funciones.
// Es similar a una "clase" en otros lenguajes de programación.

//* Persona es una `estructura` que define una entidad con Nombre y Edad.
type Persona struct {
	Nombre string
	Edad   uint8
}

func main() {

	//* Creación de una variable "maria" de tipo struct con los valores cero para cada campo.
	// Los valores cero para cada tipo de dato son: "" (cadena vacía) para strings y 0 para números.
	var maria Persona
	fmt.Println("maria:", maria)

	// Creación de una variable e instanciación de una Persona con inicialización de todos sus campos.
	// Es recomendable inicializar todos los campos especificando sus nombres para mayor claridad y seguridad.
	luis := Persona{
		Nombre: "Luis",
		Edad:   25,
	}

	fmt.Printf("Persona1: %+v\n", luis)

	//* Definición de un campo - Los campos no especificados tomarán el valor cero de su tipo de dato correspondiente.
	// No es obligatorio inicializar todos los campos, pero es recomendable hacerlo.
	andres := Persona{
		Nombre: "Andres",
	}

	fmt.Println("Persona2:", andres)

	//* Asignación de valores en orden sin necesidad de especificar el nombre de los campos.
	// Aunque no es obligatorio especificar los nombres de los campos, hacerlo aumenta la claridad del código.
	mayer := Persona{"Mayer", 23}

	fmt.Println("Persona3:", mayer)

	//* Uso del formato campo:valor para inicializar los campos en cualquier orden.
	// Es una práctica más segura y clara, especialmente cuando hay muchos campos en la estructura.
	lucas := Persona{Edad: 40, Nombre: "Lucas"}

	fmt.Println("Persona4:", lucas)

	//* Acceso a los campos de la estructura utilizando el operador punto (.).
	fmt.Println("Nombre de Luis:", luis.Nombre)
	fmt.Println("Edad de Luis:", luis.Edad)

	//* Definición de un struct anónimo: se define y se instancia en el mismo lugar sin un nombre específico.
	// Son útiles cuando solo se usan en un contexto específico y no es necesario reutilizarlos.
	objeto := struct {
		A bool
		B int
		C bool
	}{
		A: true,
		B: 42,
		C: false,
	}

	fmt.Printf("Struct anónimo: %+v\n", objeto)

	//* Uso del operador & para obtener un puntero a un struct.
	// Los cambios realizados a través del puntero afectarán al struct original,
	//ya que comparten la misma dirección de memoria.
	objetoCopy := &objeto
	objetoCopy.B = 24
	fmt.Printf("Struct original modificado a través del puntero: %+v\n", objeto)

	//* Comparación de structs.
	// En Go, los structs se pueden comparar directamente si sus campos son comparables.
	otroLuis := Persona{Nombre: "Luis", Edad: 25}
	fmt.Println("¿Son luis y otroLuis iguales?", luis == otroLuis)

	//* Uso de structs con slices y maps.
	// Los structs se pueden utilizar como elementos de slices y valores en maps.
	personas := []Persona{luis, andres, mayer}
	fmt.Println("Slice de personas:", personas)

	personasMap := map[string]Persona{
		"luis":   luis,
		"andres": andres,
		"mayer":  mayer,
	}
	fmt.Println("Map de personas:", personasMap)
}
