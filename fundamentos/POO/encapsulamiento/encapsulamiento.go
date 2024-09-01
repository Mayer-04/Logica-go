package main

import "fmt"

/*
* Encapsulamiento en Go
En Go, el concepto de encapsulamiento se logra utilizando la visibilidad de los identificadores,
tanto en los campos de las estructuras como en los métodos.

* Campos:
- Si un campo comienza con una letra mayúscula, es público (exportado) y puede ser accedido desde otros paquetes.
- Si un campo comienza con una letra minúscula, es privado (no exportado) y solo puede ser accedido desde el paquete donde fue definido.

* Métodos:
- Si un método comienza con una letra mayúscula, es público y puede ser llamado desde otros paquetes.
- Si un método comienza con una letra minúscula, es privado y solo puede ser llamado desde el paquete donde fue definido.
*/

// La estructura `Person` es pública porque su nombre comienza con la primera letra en mayúscula.
// Esto significa que puede ser utilizada en otros paquetes.
type Person struct {
	Name    string   // Campo público: puede ser accedido desde otros paquetes.
	Age     int      // Campo público: puede ser accedido desde otros paquetes.
	Live    bool     // Campo público: puede ser accedido desde otros paquetes.
	hobbies []string // Campo privado: solo puede ser accedido dentro del paquete donde se definió.
}

// GreetPerson es un método público que imprime un saludo con el nombre de la persona.
func (p Person) GreetPerson() {
	fmt.Println("Hello, my name is", p.Name)
}

// getAge es un método privado que retorna la edad de la persona.
// Solo puede ser llamado dentro del paquete donde se definió.
func (p Person) getAge() int {
	return p.Age
}

func main() {

	// Creamos una nueva instancia de `Person`.
	// En este caso,`hobbies` no puede ser accedido fuera del paquete, pero puede ser inicializado dentro de `main`.
	andres := Person{Name: "Andres", Age: 24, Live: true, hobbies: []string{"programar", "leer"}}

	// Imprimimos los valores de la instancia.
	// Los campos públicos se mostrarán, pero `hobbies` no porque es privado.
	// Este comportamiento es distinto si imprimimos `andres` con `fmt.Printf`.
	fmt.Println("andres:", andres) // Output: andres: {Andres 24 true [programar leer]}

	// Llamamos al método público `GreetPerson`.
	andres.GreetPerson() // Output: Hello, my name is Andres

	// Intentamos llamar al método privado `getAge`.
	// Aunque es privado, como estamos dentro del mismo paquete, esto funcionará.
	age := andres.getAge()
	fmt.Printf("Andres tiene %d años.\n", age) // Output: Andres tiene 24 años.
}
