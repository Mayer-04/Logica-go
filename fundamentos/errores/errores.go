package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Creando un nuevo valor de error con un mensaje personalizado
var errorDivide = errors.New("cannot divide by zero")

// TODO: El error se suele enviar como ultimo argumento de retorno en las funciones.
// TODO: En Go se utilizan valores de error explícitos en vez de excepciones.

func main() {

	data := "10"
	val, err := strconv.Atoi(data)

	// Para controlar los errores usamos el if
	// TODO: El valor 0 de los errores es "nil"
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(val)

	// Usando el nuevo valor de error creado - errorDivide
	result, err := divide(6, 3)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}

// Función que devuelve el error personalizado si b es igual a 0
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errorDivide
	}
	return a / b, nil
}
