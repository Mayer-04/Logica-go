package main

import "fmt"

/* Tenemos un objeto que almacena los salarios de nuestro equipo:

let salaries = {
  John: 100,
  Ann: 160,
  Pete: 130,
};

Escribe el código para sumar todos los salarios y almacenar el resultado en la variable sum.
En el ejemplo de arriba nos debería dar 390.

Si salaries está vacío entonces el resultado será 0.
*/

func main() {
	salaries := map[string]int{
		"John": 100,
		"Ann":  160,
		"Pete": 130,
	}

	result := SumSalaries(salaries)
	fmt.Println(result)
}

func SumSalaries(salaries map[string]int) int {
	sumaTotal := 0

	for _, salary := range salaries {
		sumaTotal += salary
	}

	return sumaTotal
}
