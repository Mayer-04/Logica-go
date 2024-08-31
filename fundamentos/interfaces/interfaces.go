package main

import "fmt"

/*
* Interfaces en Go:
Las interfaces en Go son una forma de definir un conjunto de métodos que un tipo debe implementar.

IMPORTANTE: Go utiliza un sistema de interfaces implícitas,
lo que significa que no es necesario declarar que un tipo "implementa" una interfaz; simplemente,
si un tipo posee todos los métodos definidos por una interfaz, entonces ese tipo satisface esa interfaz.

- Las interfaces pueden ser implementadas por cualquier tipo.
- Un tipo puede implementar múltiples interfaces.
- En Go, es común que las interfaces tengan un solo método.
Esto se conoce como interfaces mínimas, lo que permite una mayor flexibilidad y reusabilidad.
- Es una buena práctica nombrar las interfaces con un sufijo `-er`,
como `Printer` en lugar de `Print`, siguiendo una convención del lenguaje.
- Las interfaces pueden ser compuestas, es decir, una interfaz puede incluir otras interfaces.
- IMPORTANTE: Puedes implementar interfaces definidas en otros paquetes de Go, no solo las que tú defines.
*/

// Printer es una interfaz mínima que define un solo método Print.
type Printer interface {
	Print()
}

// Logger es una interfaz que define un solo método Log.
type Logger interface {
	Log(message string)
}

// Person es una estructura que contiene el nombre y la edad de una persona.
type Person struct {
	Name string
	Age  int
}

// PrinterLogger es una interfaz compuesta que incluye las interfaces Printer y Logger.
// Cualquier tipo que implemente ambos métodos (Print y Log) satisfará la interfaz PrinterLogger.
type PrinterLogger interface {
	Printer
	Logger
}

// Print es un método de Person que cumple con la interfaz `Printer`.
// Esto significa que cualquier instancia de Person puede ser asignada a una variable de tipo Printer.
func (p Person) Print() {
	fmt.Printf("Hola, mi nombre es %s y tengo %d años.\n", p.Name, p.Age)
}

// Log es un método de Person que cumple con la interfaz Logger.
// Esto significa que cualquier instancia de Person puede ser asignada a una variable de tipo Logger.
func (p Person) Log(message string) {
	fmt.Printf("[LOG]: %s\n", message)
}

func main() {

	// Crear una instancia de Person y asignarla a una variable de tipo Printer.
	// Esto es posible porque Person implementa el método Print, que es el único
	// requisito para satisfacer la interfaz Printer.
	var p Printer = Person{
		Name: "Mayer",
		Age:  24,
	}
	p.Print() // Hola, mi nombre es Mayer y tengo 24 años.

	// Crear otra instancia de Person y usar el método Print directamente.
	// Aquí no estamos usando la interfaz, sino la estructura directamente.
	// El método `Print` de Person ya implementa implícitamente la interfaz Printer.
	robert := Person{
		Name: "Robert",
		Age:  59,
	}
	robert.Print() // Hola, mi nombre es Robert y tengo 59 años.

	// Crear una instancia de Person y asignarla a una variable de tipo Logger.
	// Esto es posible porque Person implementa el método Log.
	var cristiano Logger = Person{
		Name: "Cristiano",
		Age:  39,
	}
	cristiano.Log("Soy el bicho...SIUUU!") // [LOG]: Soy el bicho...SIUUU!

	// Crear una instancia de Person y asignarla a una variable de tipo PrinterLogger.
	// Esto es posible porque Person implementa ambos métodos: Print y Log,
	// cumpliendo así con la interfaz compuesta PrinterLogger.
	var person PrinterLogger = Person{
		Name: "Messi",
		Age:  37,
	}
	person.Print()                 // Hola, mi nombre es Messi y tengo 37 años.
	person.Log("Hola, soy Messi!") // [LOG]: Hola, soy Messi!
}
