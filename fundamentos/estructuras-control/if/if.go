package main

import "fmt"

/*
* Estructura if: Si
Si la condición es verdadera, se ejecutará el bloque de código del if.

- En Go no existe el operador ternario como en otros lenguajes de programación.
- No hay conversión automática de valores a booleano (por ejemplo, 0 no es false).
- La condición dentro de un if debe ser `explícitamente` una expresión booleana.
- No se permite la asignación dentro de la condición; hacerlo causará un error de compilación.
*/

func main() {

	// La condición no debe estar entre paréntesis.
	// La condición debe ser "true" para que se ejecute el bloque de código.
	if true {
		fmt.Println("¡Hola Mundo!")
	}

	edad := 18

	// * Estructura if-else.
	// Si la condición (edad >= 18) es verdadera, se ejecutará el bloque de código del if.
	// Si la condición es falsa, se ejecutará el bloque de código del else.
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

	// * If con declaración de variable corta:
	// La variable 'mayor' se declara en la condición y su ámbito es el bloque if-else.
	// Esta estructura permite realizar una comprobación en una sola línea.
	if mayor := edad >= 18; mayor {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

	// * Estructura if-else-if:
	// Permite evaluar múltiples condiciones en secuencia.
	// La primera condición verdadera será la que determine el bloque de código a ejecutar.
	// Si ninguna condición es verdadera, se ejecutará el bloque else final (si existe).
	if edad < 13 {
		fmt.Println("Eres un niño")
	} else if edad < 18 {
		fmt.Println("Eres un adolescente")
	} else {
		fmt.Println("Eres un adulto")
	}

	// * Aserción de tipo (Type Assertion).
	// Es una operación que permite obtener el valor oculto de un valor almacenado en una variable de tipo interface,
	// es decir, una interfaz vacía que acepta cualquier tipo de dato.
	// "Oye, sé que esta variable guarda algo, y creo que es un tipo específico.
	// Quiero que me lo muestres como ese tipo".
	var i interface{} = 10

	if v, ok := i.(int); ok {
		fmt.Println("i es un int:", v)
	} else {
		fmt.Println("i no es un int")
	}

}
