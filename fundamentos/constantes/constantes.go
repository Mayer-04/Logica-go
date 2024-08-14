package main

import "fmt"

// * IMPORTANTE: Las constantes no necesitan ser usadas en su totalidad para compilar el código.
// * Declarar una constante sin usarla no provocará un mensaje de error.
// * Las constantes se deben declarar y asignar inmediatamente un valor.
// * Es recomendable declararlas a nivel de paquete cuando se necesita un valor constante en múltiples lugares.

// Constante exportable a nivel de paquete
const Pi = 3.1416

// Constante no exportable, solo visible dentro del paquete
const version = "1.0.0"

// Constante a nivel de paquete, no exportable
const animal = "🐯"

// * Uso de iota para la creación de secuencias de valores incrementales en declaraciones const
// * iota es una característica incorporada en Go que comienza en 0 y se incrementa automáticamente.
// * Ideal para enumeraciones o listas secuenciales.
const (
	Lunes     = iota // Lunes == 0
	Martes           // Martes == 1 (iota se incrementa automáticamente)
	Miércoles        // Miércoles == 2 (iota se incrementa automáticamente)
	Jueves           // Jueves == 3 (iota se incrementa automáticamente)
	Viernes          // Viernes == 4 (iota se incrementa automáticamente)
	Sábado           // Sábado == 5 (iota se incrementa automáticamente)
	Domingo          // Domingo == 6 (iota se incrementa automáticamente)
)

func main() {
	// Constantes
	// * Las constantes deben declararse con la palabra clave 'const' seguida del nombre y tipo de la constante.
	// * No podemos utilizar el operador de asignación corta (:=) para definir una constante.
	const gender string = "Masculino"
	fmt.Println("género:", gender)

	// Agrupación de constantes
	// * Es posible agrupar constantes relacionadas en un solo bloque.
	const (
		fruit1 = "🍎"
		fruit2 = "🍐"
	)

	fmt.Println(fruit1, fruit2)

	// Imprimir constantes definidas a nivel de paquete
	fmt.Println("animal:", animal)
	fmt.Println("versión:", version)
}
