package main

import (
	"fmt"
	"unique"
)

/*
* Paquete unique: Unico
El paquete `unique` es una nueva característica de Go que se integró a partir de la versión 1.23.
- Se enfoca en la optimización de la memoria y la eficiencia, aprovechando una técnica conocida como "interning".
- Ahorro de memoria: Al almacenar una única copia de los valores repetidos, se reduce significativamente el uso de memoria.
- Se utiliza con cualquier valor comparable.
- Se ocupa de la concurrencia de manera automática (como lo hace con `sync.Map`).
- Evita el crecimiento descontrolado de memoria mediante una limpieza automática integrada con el garbage collector (GC) de Go.
- En lugar de comparar los bytes de los valores, compara sus "handles" (punteros), lo que resulta en comparaciones mucho más rápidas.
- El paquete es seguro para ser usado con múltiples goroutines.
- Cuando un valor internado ya no es necesario, el garbage collector lo elimina automaticamente.

* Interning: Internamiento
El interning es una técnica que garantiza que solo se guarde una copia de un valor repetido en memoria.
Si el mismo valor se solicita varias veces, en lugar de crear nuevas copias, se reutiliza la copia existente.
*/

func main() {
	// Internamiento de cadenas (string)
	name := unique.Make("Mayer")  // Interna el valor "Mayer"
	name2 := unique.Make("Mayer") // Dado que "Mayer" ya existe, name2 apunta a la misma ubicación en memoria

	// Verificamos que ambos apuntan al mismo "handle" (la referencia interna es la misma)
	// Las comparaciones entre "handles" son eficientes porque comparamos las direcciones de memoria.
	fmt.Println("name == name2:", name == name2) // true, ambas referencias son iguales

	// Internamiento de otra cadena (string)
	lastname := unique.Make("Prada")                   // Nuevo valor "Prada", distinta referencia
	fmt.Println("name == lastname:", name == lastname) // false, porque son valores diferentes

	// Internamiento de números (int)
	edad := unique.Make(24)                      // Interna el valor 24
	edad2 := unique.Make(24)                     // Reutiliza la referencia de 24
	fmt.Println("edad == edad2:", edad == edad2) // true, ambos apuntan al mismo valor internado
}
