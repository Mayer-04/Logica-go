package main

import "fmt"

/*
* Constructores en Go
Go no tiene constructores como en otros lenguajes de programación orientado a objetos.
En su lugar, se utilizan funciones que retornan una instancia de una estructura (struct).

- Un "constructor" es una función que se utiliza para crear e inicializar una nueva instancia de una estructura.
- Por convención, estas funciones suelen llamarse `New` o `NewNombreStruct`.
- Las funciones constructoras pueden incluir validaciones para asegurar que los valores iniciales son correctos.
- Estas funciones pueden devolver punteros o valores según la necesidad.
*/

// Definimos una estructura llamada `Persona`
type Persona struct {
	Nombre string
	Edad   int
}

//* New crea una nueva instancia de `Persona`.
// Recibe el nombre y la edad como parámetros y retorna un puntero a una nueva instancia de `Persona`.
func New(nombre string, edad int) *Persona {
	// Creamos y retornamos un puntero a `Persona` con los valores iniciales provistos.
	return &Persona{
		Nombre: nombre,
		Edad:   edad,
	}
}

//* NewPersona crea una nueva instancia de `Persona` con validaciones.
// Si el nombre es vacío o la edad es negativa, retorna un error.
func NewPersona(nombre string, edad int) (*Persona, error) {

	// Validación para nombre vacío
	if nombre == "" {
		return nil, fmt.Errorf("el nombre no puede estar vacío")
	}

	// Validación para edad negativa
	if edad < 0 {
		return nil, fmt.Errorf("la edad no puede ser negativa")
	}

	// Si los datos son válidos, retorna un puntero a `Persona`.
	return &Persona{
		Nombre: nombre,
		Edad:   edad,
	}, nil
}

func main() {
	// Crear una nueva instancia de Persona usando el constructor básico `New`.
	persona := New("Mayer", 24)
	fmt.Printf("persona: %+v\n", persona)                              // Output: persona: &{Nombre:Mayer Edad:24}
	fmt.Printf("nombre: %s, edad: %d\n", persona.Nombre, persona.Edad) // Output: nombre: Mayer, edad: 24

	// Crear una nueva instancia de Persona usando el constructor `NewPersona` con validaciones.
	persona2, err := NewPersona("Andres", -24)
	if err != nil {
		// En caso de error, se imprime el mensaje de error.
		fmt.Println(err) // Output: la edad no puede ser negativa
		return
	}

	// Si no hubo errores, se imprimen los valores de `persona2`.
	fmt.Printf("persona2: %+v\n", persona2)
	fmt.Printf("nombre: %s, edad: %d\n", persona2.Nombre, persona2.Edad)
}
