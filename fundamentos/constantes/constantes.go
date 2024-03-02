package main

import "fmt"

// TODO: IMPORTANTE - Las constantes no necesitan ser usadas en su totalidad para compilar el código

// Constante a nivel de paquete
const animal = "🐯"

// * Creación de secuencias de valores incrementales dentro de las declaraciones const
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
	const GENDER string = "Masculino"

	fmt.Println("género:", GENDER)

	// Agrupar constantes
	const (
		fruit1 = "🍎"
		fruit2 = "🍐"
	)

	fmt.Println(fruit1, fruit2)

	fmt.Println(animal)

}
