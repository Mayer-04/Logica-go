package main

import "fmt"

/* Para evitar problemas a la hora de AGREGAR elementos a una "map", asegúrese de crear un "map" vacío
(no una asignación nil) mediante la función make, esta regla solo se aplica al agregar elementos */

//!NOTE: El valor 0 de un map es nil si no se ha inicializa utilizando make()

func main() {

	// Creando un map literal - declarar e inicializar inmediatamente el valor
	// TODO: Obligatoria la coma al final de cada elemento
	var edades = map[string]int{
		"Mayer":  24,
		"Andres": 23,
		"Luis":   25,
		"Maria":  26,
	}
	fmt.Println("Edades:", edades)

	// Creando un map con valores iniciales
	nuevoMapa := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("nuevoMapa:", nuevoMapa)

	// Inicializando un mapa vacío
	mapaVacio := map[int]string{}
	fmt.Println("mapa Vacío:", mapaVacio) // Output: mapa Vacío: map[]

	// Creando e inicializando un map con make
	//* Solo podemos agregar la "capacidad" como segundo argumento al crear el map con make()
	edades2 := make(map[string]int, 2)

	edades2["Mayer"] = 24
	edades2["Andres"] = 23

	fmt.Println("Edades:", edades2)

	// Accediendo a un elemento del map
	// TODO: Si no existe la llave, nos devuelve el valor cero del tipo de dato correspondiente
	fmt.Println(nuevoMapa["a"])
	fmt.Println("Llave d:", nuevoMapa["d"]) // Output: Llave d: 0

	//* Eliminar un elemento del map a partir de su llave(clave)
	delete(nuevoMapa, "a")

	// Asignación de dos valores - Verifica si un elemento existe en el map
	//* "value" es el valor de la clave "a", si la clave no existe "value" devuelve el valor cero
	//* "ok" es un booleano que indica si existe o no
	value, ok := nuevoMapa["a"]
	fmt.Printf("Valor: %v - Existe: %t\n", value, ok)

	// TODO: Con la función incorporada "clear()" podemos eliminar todos los elementos, lo que da un mapa vacío
	clear(edades)
	fmt.Println(edades)
}
