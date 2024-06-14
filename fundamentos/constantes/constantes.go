package main

import "fmt"

//* IMPORTANTE: Las constantes no necesitan ser usadas en su totalidad para compilar el código
//* Comenzar las constantes en Go con una letra mayúscula si quieres que sean visibles fuera del paquete
//* Declarar una constante sin usarla no provocara un mensaje de error

// Exportable, visible fuera del paquete en el que están definidas.
const Pi = 3.1416

// No exportable, solo visible dentro del paquete
const version = "1.0.0"

// Constante a nivel de paquete
const animal = "🐯"

//* iota: Creación de secuencias de valores incrementales dentro de las declaraciones const
// TODO: iota siempre empieza en 0
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
	// TODO: Se debe declarar y asignar inmediatamente el valor
	// TODO: No podemos utilizar el operador de variable corta :=
	// TODO: Recomandable usarlas a nivel de paquete
	const gender string = "Masculino"
	fmt.Println("género:", gender)

	// Agrupar constantes
	const (
		fruit1 = "🍎"
		fruit2 = "🍐"
	)

	fmt.Println(fruit1, fruit2)

	fmt.Println("animal:", animal)
	fmt.Println("versión:", version)
}
