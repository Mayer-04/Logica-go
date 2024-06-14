package main

import "fmt"

type Person struct {
	Name     string
	LastName string
	Age      int
	Hoobies  []string
}

// TODO: new(T) - Crear una instancia de un tipo de dato concreto y devuelve un puntero a ese tipo de dato

func main() {

	var x = new(int)

	fmt.Println("Valor de x en memoria:", x) // Output: Valor de x en memoria: 0xc00000a0d8

	*x = 10

	//  Acceder al valor almacenado en la dirección de memoria con el operador de desreferenciación
	fmt.Println("Nuevo valor de x:", *x)

	// Go inicializa automáticamente los campos del struct con valores cero de su tipo de dato correspondiente
	var person = new(Person)
	fmt.Println(person) // Output: &{"" "" 0 []}

	// Asignando valores a los campos del struct
	person.Name = "Mayer"
	person.LastName = "Chaves"
	person.Age = 23
	person.Hoobies = []string{"Programar", "Leer"}

	// Imprimiendo el struct con el operador de desreferenciación para acceder a sus campos
	fmt.Println(*person)
}
