package main

import (
	"fmt"
	"strings"
)

// TODO: Las funciones en Go son de primera clase
// TODO: Los parámetros de funciones son por valor no por referencia - Se trabaja con la copia del valor
func main() {

	saludar("Mayer", "Chaves")

	// Ejemplo con función por referencia
	num := 5
	// Obtenemos la dirección en memoria de la variable "num"
	increment(&num)

	fmt.Println(num) // El valor de num se cambia directamente

	// TODO: RETURN
	result := retornar(1, 3)
	fmt.Println(result)

	// TODO: Multiples valores a retornar
	// Podemos ignorar u omitir algun valor con el operador blank
	lower, upper := multiplesRetornos("Mayer")
	fmt.Println(lower, upper)

	// TODO: Función variática
	fmt.Println(variatica(1, 2, 3, 4, 5))

	// TODO: Función anonima
	anonima()

	// TODO: Función anonima autoinvocada
	func(name string) {
		fmt.Println("Hello", name)
	}("mayer")

	// TODO: Segunda manera de las funciones anonimas con variable corta
	ejemplo := func() {
		fmt.Println("Soy una función anonima corta")
	}

	ejemplo()
}

// TODO: Características de las funciones ----------------------------------------------------------------------------

// Agrupar parámetros donde los dos sean del mismo tipo de dato
func saludar(name, lastname string) {
	fmt.Printf("Nombre %s con apellido %s", name, lastname)
}

// Para trabajar con parámetros por referencia se deben utilizar "Punteros"
func increment(x *int) {
	// Operador desreferenciación - Acceder al valor almacenado en la dirección de memoria
	*x++
}

// Función con valor de retorno - return
func retornar(num1, num2 int) int {
	return num1 + num2
}

// Funciones con multiples retornos
// TODO: Se debe incluir entre paréntesis los multiples valores a retornar
// Podemos nombrar los parámetros de retorno - NO RECOMENDADO EN FUNCIONES COMPLEJAS
func multiplesRetornos(text string) (string, string) {

	lower := strings.ToLower(text)
	upper := strings.ToUpper(text)

	return lower, upper
}

// Funciones variáticas
// TODO: Pasar un número indefinido de argumentos - Anteponer al tipo de datos tres puntos "..."
func variatica(nums ...int) int {

	// el parámetro nums se convierte en un "slice"

	var total int

	for _, num := range nums {
		total += num
	}

	return total

}

// TODO: Funciones anonimas - No tienen un nombre de función
var anonima = func() {
	fmt.Println("Soy una función anonima")
}
