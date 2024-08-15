package main

import (
	"fmt"
	"strings"
)

func main() {
	var cadena string = "ANACONDA"
	var caracter string = "u"

	result := NumeroDeVeces(cadena, caracter)

	fmt.Println(result)
}

/*
Escribe una función que tome una cadena de texto y un carácter específico,
y devuelva el número de veces que ese carácter aparece en la cadena.
*/
func NumeroDeVeces(cadena, caracter string) int {

	var contador int = 0

	newString := strings.ToLower(cadena)
	newCharacter := strings.ToLower(caracter)

	for _, value := range newString {
		if string(value) == newCharacter {
			contador++
		}
	}

	return contador

}
