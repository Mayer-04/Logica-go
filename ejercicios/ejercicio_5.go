package ejercicios

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// Dado un número entero x, devuelve verdadero si x es un palíndromo y falso en caso contrario.

func Ejercicio5() {

	var number = 121

	result := isPalindrome(number)

	fmt.Println(result)
}

func isPalindrome(x int) bool {
	var nuevoString string = strconv.Itoa(x)
	arregloStrings := strings.Split(nuevoString, "")

	slices.Reverse(arregloStrings)

	var unir string = strings.Join(arregloStrings, "")

	nuevoNumber, err := strconv.Atoi(unir)

	if err != nil {
		fmt.Println(err)
	}

	return x == nuevoNumber

}
