package main

import (
	"fmt"
	"reflect"
)

/*
Tienes una función que recibe un objeto como parámetro.
La función debe retornar un array con el nombre de las propiedades que su tipo sea boolean.

Por ejemplo, para el objeto { a: true, b: 42, c: false } la función debe retornar ['a', 'c'],
ya que son las dos propiedades que tienen valores booleanos.
*/

type Map map[string]any

func main() {
	obj := Map{"a": true, "b": 42, "c": false}
	result := propiedadesBooleanas(obj)
	fmt.Println(result)
}

func propiedadesBooleanas(obj map[string]any) []string {
	var booleanProperties []string

	for key, value := range obj {
		typeof := reflect.TypeOf(value).Kind()

		if typeof == reflect.Bool {
			booleanProperties = append(booleanProperties, key)
		}
	}

	return booleanProperties
}
