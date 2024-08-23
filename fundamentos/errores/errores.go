package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
* En Go se utilizan valores de error explícitos en vez de excepciones.
- Los errores se suelen enviar como último argumento de retorno en las funciones.
- Para controlar los errores usamos la estructura de control `if`.
- El valor 0 de los errores es `nil`.
*/

// La función 'New()' del paquete "errors" crea un nuevo valor de error con un mensaje personalizado
var errorDivide = errors.New("cannot divide by zero")

func main() {

	data := "10"
	val, err := strconv.Atoi(data)

	// Controlando el error con 'if'
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)

	// Usando el nuevo valor de error creado - errorDivide
	result, err := divide(6, 0)

	if err != nil {
		fmt.Println("error:", err) // Error: cannot divide by zero
		return
	}
	fmt.Println("result:", result)
}

// Función que devuelve el error personalizado si b es igual a 0
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errorDivide
	}

	// Retorna el resultado de la division y "nil" haciendo referencia a que no hay error
	return a / b, nil
}
