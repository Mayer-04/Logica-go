package main

import "fmt"

func main() {

	// Creando un map literal - declarar y asignar inmediatamente el valor
	// TODO: Obligatoria la coma al final de cada elemento
	var edades = map[string]int{
		"Mayer":  24,
		"Andres": 23,
		"Luis":   25,
		"Maria":  26,
	}

	fmt.Println("Edades:", edades)

	// Creando e inicializando un map con make
	edades2 := make(map[string]int)

	edades2["Mayer"] = 24
	edades2["Andres"] = 23

	fmt.Println("Edades:", edades2)

	// Creando un map con valores iniciales
	nuevoMapa := map[string]int{"a": 1, "b": 2, "c": 3}

	fmt.Println(nuevoMapa)

	// Accediendo a un elemento del map
	// TODO: Si no existe la llave, nos devuelve el valor 0 del tipo de dato correspondiente
	fmt.Println(nuevoMapa["a"])
	fmt.Println("Llave d:", nuevoMapa["d"]) // Output: Llave d: 0

	// Eliminar un elemento del map a partir de su llave(clave) - delete
	delete(nuevoMapa, "a")

	// Verificar si un elemento existe en el map
	content, ok := nuevoMapa["a"]

	fmt.Printf("Contenido: %v - Existe: %t", content, ok)
}
