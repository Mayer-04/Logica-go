package main

import "fmt"

/*
* Los maps en Go son estructuras de datos no ordenadas que almacenan pares clave-valor.
- Una clave puede ser de cualquier tipo comparable (int, uint, string, struct, etc).
- Un valor puede ser de cualquier tipo.
- Los maps proporcionan acceso rápido a los datos y son muy útiles para almacenar conjuntos de datos
donde se necesita buscar elementos por clave.

IMPORTANTE:
- Para evitar problemas al agregar elementos a un map, asegúrese de crearlo usando la función `make()`
para inicializarlo correctamente.
- El valor cero de un map no inicializado es `nil`, lo que significa que no apunta a ninguna estructura de datos.
*/

func main() {

	// Creando un map literal - declarando e inicializando el map con valores predeterminados
	// NOTA: Es obligatorio incluir una coma al final de cada elemento en un map literal.
	var edades = map[string]int{
		"Mayer":  24,
		"Andres": 23,
		"Luis":   25,
		"Maria":  26,
	}
	fmt.Println("Edades:", edades)

	// Creando un map con valores iniciales en una sola línea
	nuevoMapa := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("nuevoMapa:", nuevoMapa)

	// Inicializando un map vacío utilizando un map literal
	// Es común usar esta técnica para evitar un nil map, el cual no permite agregar elementos.
	mapaVacio := map[int]string{}
	fmt.Println("Mapa vacío:", mapaVacio) // Output: Mapa vacío: map[]

	// Creando e inicializando un map vacío utilizando `make`
	// Podemos especificar la capacidad inicial como segundo argumento del map para optimizar el rendimiento
	edades2 := make(map[string]int, 2)

	// Agregando elementos al map
	edades2["Mayer"] = 24
	edades2["Andres"] = 23

	fmt.Println("Edades2:", edades2)

	// Accediendo a un elemento del map
	// Si la clave no existe, se devuelve el valor cero del tipo de dato correspondiente.
	fmt.Println("Valor de la llave 'a':", nuevoMapa["a"])
	fmt.Println("Valor de la llave 'd':", nuevoMapa["d"]) // Output: Valor de la llave 'd': 0

	// Eliminando un elemento del map utilizando su clave
	delete(nuevoMapa, "a")

	// Verificando si una clave existe en el map y obteniendo su valor
	// "value" es el valor de la clave "a". Si la clave no existe, "value" será el valor cero.
	// "ok" es un booleano que indica si la clave existe en el map.
	value, ok := nuevoMapa["a"]
	fmt.Printf("Valor: %v - Existe: %t\n", value, ok)

	// Limpiando el map eliminando todos sus elementos usando la función `clear()`
	// `clear()` es una función moderna que simplifica la eliminación de todos los elementos de un map.
	clear(edades)
	fmt.Println("Edades después de clear():", edades)

	// Recorriendo un map con un bucle `for range`
	// El orden de iteración de los elementos en un map es aleatorio y puede cambiar entre ejecuciones del programa.
	for nombre, edad := range edades2 {
		fmt.Printf("Nombre: %q, Edad: %d\n", nombre, edad)
	}
}
