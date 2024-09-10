package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
* Manejo de Errores en Go
En Go, los errores se tratan de manera `explícita` en lugar de usar excepciones, como en otros lenguajes.
Esto proporciona un manejo de errores más claro y controlado. Algunos puntos clave son:
- Los errores se devuelven como el último valor de retorno en las funciones.
- Para manejar los errores, generalmente usamos la estructura de control `if`.
- El valor por defecto de un error es `nil`, lo que indica que no hubo error.
- No se deben ignorar los errores devueltos, incluso por operaciones con `defer`,
si deseas ignorar un error, utiliza el operador blank `_`.

* Es un error común manejar un mismo error varias veces. Debemos recordar que:
- Manejar un error significa registrarlo (log) o devolverlo (return), no ambas cosas.
- Si manejas un error, es importante hacerlo solo una vez.
- Para envolver un error y conservar la causa original, se puede usar `fmt.Errorf` con la directiva `%w`.
La comparación de este tipo de errores se debe hacer con `errors.As`.
- Para comparar un error con un valor específico, se debe usar `errors.Is` en lugar de `==`.
- Si no deseas exponer el error original a quien llama a la función,
no envuelvas el error usando `%w`; en su lugar, usa `%v` para solo formatear el mensaje.
*/

// La función 'New()' del paquete "errors" crea un nuevo valor de error con un mensaje personalizado.
var errorDivide = errors.New("cannot divide by zero")

func main() {
	// Convertimos un string a un entero utilizando 'strconv.Atoi'.
	// Si falla, se devuelve un error.
	value, err := strconv.Atoi("10")
	if err != nil {
		// Imprimimos un mensaje de error y nos salimos del programa.
		fmt.Printf("cannot convert %q to int: %v\n", "10", err)
		return
	}

	// Imprimimos el valor convertido.
	fmt.Println("Valor convertido:", value)

	// Llamamos a la función 'divide' que nos devuelve un error.
	result, err := divide(6, 0)
	// Comprobamos si hubo un error.
	if err != nil {
		// Si el error es específico de división, lo imprimimos.
		if errors.Is(err, errorDivide) {
			fmt.Println("Error en la división:", err)
		}
		// Nos salimos del programa.
		return
	}
	// Imprimimos el resultado.
	fmt.Println("Resultado:", result)
}

// Función que realiza una división y devuelve un error si b es igual a 0.
func divide(a, b float64) (float64, error) {
	// Verificamos si el divisor es cero.
	if b == 0 {
		// Devolvemos un valor de error personalizado.
		return 0, errorDivide
	}
	// Si no hay error, devolvemos el resultado de la división y nil como indicación de que no hubo errores.
	return a / b, nil
}
