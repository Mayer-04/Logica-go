package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
* Strings: Cadenas de caracteres
- Los strings son una colección de caracteres.
- Los strings son inmutables, es decir, no se pueden modificar.
- Los strings se declaran entre comillas dobles.
- Los strings se pueden concatenar con cualquier otro tipo de dato con la función incorporada `fmt.Sprintf()`.
- Para escribir un salto de línea, se escribe `\n`.
- Si escribes una letra en comillas simples, se interpreta como un tipo `rune`, alias para el tipo int32.
*/

func main() {

	//* Los strings se declaran entre comillas dobles.
	name := "Mayer"
	fmt.Printf("tipo: %T - valor: %s\n", name, name)

	// Salto de línea `\n`.
	secondName := "Andres"
	fmt.Printf("tipo: %T - valor: %v\n", secondName, secondName)

	// Interpolación.
	// Se utiliza el paquete "fmt" con su función "Sprintf" para formatear una cadena con valores.
	fullName := fmt.Sprintf("Mi nombre es %s %s", name, secondName)
	fmt.Println("interpolación:", fullName)

	//* Paquete "strings" - Contiene métodos para trabajar con cadenas de texto

	// Longitud de un string con la función incorporada "len()".
	fmt.Println("longitud de name:", len(name))

	// Concatenar strings.
	nameAndSecondName := name + " " + secondName
	fmt.Println("concatenar strings:", nameAndSecondName)

	// Verifica si la cadena contiene una subcadena - Retorna un booleano.
	str := strings.Contains(name, "Mayer")
	fmt.Println(str)

	// Verifica el índice de la subcadena - Retorna un entero
	str2 := strings.Index(name, "y")
	fmt.Println(str2)

	// Convierte una cadena a minúsculas.
	str3 := strings.ToLower(name)
	fmt.Println(str3)

	// Reemplaza todas las ocurrencias de una subcadena por otra.
	str4 := strings.Replace(name, "Mayer", "Chaves", -1)
	fmt.Println(str4)

	// Divide una cadena en partes de acuerdo a un separador - Retorna un slice de strings.
	str5 := strings.Split(name, "")
	fmt.Println(str5)

	// Verifica si una cadena termina con una subcadena - Retorna un booleano.
	str6 := strings.HasSuffix(name, "er")
	fmt.Println(str6)

	// Eliminar los espacios del inicio y final.
	str7 := strings.TrimSpace(name)
	fmt.Println(str7)

	//* Paquete "strconv" - Contiene métodos para convertir strings a números enteros o flotantes

	// Convertir un string a un entero.
	age := "23"
	// Omitimos el error por temas didacticos.
	age2, _ := strconv.Atoi(age)
	fmt.Printf("tipo: %T - valor: %v\n", age2, age2)

	// Convertir un string que representa un número a un float64.
	weight := "70.5"
	weight2, _ := strconv.ParseFloat(weight, 64)
	fmt.Printf("tipo: %T - valor: %v", weight2, weight2)

}
