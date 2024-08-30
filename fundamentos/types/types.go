package main

import "fmt"

/*
* Types: Tipos en Go
Los tipos en Go son una forma de definir la estructura y el comportamiento de los datos.

- Pueden ser tipos básicos, como int, string, etc., o tipos definidos por el usuario,
como estructuras (struct), alias de tipos existentes, y funciones.
*/

// Creando un tipo Person que es una estructura.
type Person struct {
	Name string
	Age  int
}

// Creando un tipo que es un slice de estructuras Person.
// Esto permite trabajar con múltiples instancias de Person como un solo tipo.
type People []Person

// Definición de un tipo personalizado basado en un tipo existente.
// MyNumber es un alias para el tipo int, lo que nos permite agregar métodos específicos.
type MyNumber int

// Creamos un método para el tipo MyNumber.
// Esto nos permite asociar comportamientos (métodos) específicos al tipo.
func (m MyNumber) String() string {
	return fmt.Sprintf("%d", m)
}

// Definiendo un tipo para una función que recibe dos strings como parámetros.
// Los tipos de función permiten definir y trabajar con funciones como valores.
type MyFunction func(string, string)

// Implementando un método para el tipo MyFunction.
// Este método 'hello' invoca la función almacenada en MyFunction.
func (m MyFunction) hello(name, lastname string) {
	m(name, lastname)
}

func main() {
	// Declaración de alias.
	// Un alias de tipo es otra forma de referirse a un tipo existente.
	type myPerson = Person
	person := myPerson{"Mayer", 24}
	// Imprimir el valor de person.
	fmt.Printf("person: %+v\n", person) // Output: person: {Name:Mayer Age:24}

	// Definimos una variable de tipo MyNumber.
	var number MyNumber = 11
	fmt.Printf("myNumber: %T\n", number)                 // Output: myNumber: main.MyNumber
	fmt.Printf("método String(): %s\n", number.String()) // Output: 11

	// Definimos una variable de tipo MyFunction y la asignamos a una función anónima.
	var function MyFunction = func(name, lastname string) {
		fmt.Printf("Hello, %s %s!\n", name, lastname)
	}
	// Usamos el método hello para llamar a la función asignada a MyFunction.
	function.hello("Mayer", "Andres") // Output: Hello, Mayer Andres!

	// Combinando tipos personalizados
	// Creamos una lista de personas (People) usando el tipo personalizado `type People []Person`.
	people := People{
		{Name: "John", Age: 30},
		{Name: "Cristiano", Age: 39},
	}
	// Imprimimos la lista de personas.
	fmt.Printf("people: %+v\n", people) // Output: people: [{Name:John Age:30} {Name:Cristiano Age:39}]

	// Alias y métodos personalizados
	// Usamos el alias myPerson y el método String() del tipo MyNumber.
	person2 := myPerson{"Messi", 37}
	number2 := MyNumber(42)
	fmt.Printf("person2: %+v\n", person2)                             // Output: person2: {Name:Messi Age:37}
	fmt.Printf("number2 (con método String): %s\n", number2.String()) // Output: 42
}