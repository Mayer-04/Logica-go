package main

import (
	"fmt"
	"reflect"
)

func main() {
	var obj = map[string]any{"a": true, "b": 42, "c": false}

	result := propiedadesBooleanas(obj)

	fmt.Println(result)
}

func propiedadesBooleanas(obj map[string]any) []string {
	/* Tienes una función que recibe un objeto como parámetro. La función debe retornar un array con el nombre de las propiedades que su tipo sea boolean.

	Por ejemplo, para el objeto { a: true, b: 42, c: false } la función debe retornar ['a', 'c'] ya que son las dos propiedades que tienen valores booleanos. */

	var booleanProperties []string

	for key, value := range obj {
		typeof := reflect.TypeOf(value).Kind()

		if typeof == reflect.Bool {
			booleanProperties = append(booleanProperties, key)
		}
	}

	return booleanProperties

}
