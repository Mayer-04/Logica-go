package main

import "fmt"

// TODO: IMPORTANTE - Las constantes no necesitan ser usadas en su totalidad para compilar el código
// TODO: Comenzar las constantes en Go con una letra mayúscula si quieres que sean visibles fuera del paquete.

// Exportable, visible fuera del paquete en el que están definidas.
const Pi = 3.1416

// No exportable, solo visible dentro del paquete
const version = "1.0.0"

// Constante a nivel de paquete
const Animal = "🐯"

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
	const Gender string = "Masculino"

	fmt.Println("género:", Gender)

	// Agrupar constantes
	const (
		Fruit1 = "🍎"
		Fruit2 = "🍐"
	)

	fmt.Println(Fruit1, Fruit2)

	fmt.Println(Animal)
	fmt.Println(version)

}
