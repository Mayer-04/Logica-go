package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: Si escribes con comillas simples una letra sera de tipo "rune", alias para el tipo int32

func main() {

	// TODO: Los strings se declaran entre comillas dobles
	var name string = "Mayer"

	fmt.Printf("Tipo: %T - Valor: %v", name, name)

	secondName := "Andres"

	//* Salto de línea - \n
	fmt.Printf("Tipo: %T\n - Valor: %v", secondName, secondName)

	// TODO:  Interpolación - fmt.Sprintf
	fullName := fmt.Sprintf("Mi nombre es %s %s", name, secondName)
	fmt.Println("Interpolación:\n", fullName)

	//* Paquete "strings" contiene métodos para trabajar con cadenas de texto ----------------------------------------

	// Longitud de un string
	fmt.Println("Longitud de name:\n", len(name))

	// Concatenar strings
	var nameAndSecondName = name + " " + secondName

	fmt.Println("\n", nameAndSecondName)

	// Verifica si la cadena contiene una subcadena - Retorna un booleano
	str := strings.Contains(name, "Mayer")

	fmt.Println(str)

	// Verifica el índice de la subcadena
	str2 := strings.Index(name, "y")

	fmt.Println(str2)

	// Convierte una cadena a minúsculas
	str3 := strings.ToLower(name)

	fmt.Println(str3)

	// Reemplaza todas las ocurrencias de una subcadena por otra
	str4 := strings.Replace(name, "Mayer", "Chaves", -1)

	fmt.Println(str4)

	// Divide una cadena en partes de acuerdo a un separador - Retorna un slice de strings
	str5 := strings.Split(name, "")

	fmt.Println(str5)

	// Verifica si una cadena termina con una subcadena - Retorna un booleano
	str6 := strings.HasSuffix(name, "er")

	fmt.Println(str6)

	// Eliminar los espacios del inicio y final
	str7 := strings.TrimSpace(name)

	fmt.Println(str7)

	//* Paquete "strconv" contiene métodos para convertir strings----------------------------------------------------

	// Convertir un string a un entero
	age := "23"
	age2, _ := strconv.Atoi(age)

	fmt.Printf("Tipo: %T - Valor: %v\n", age2, age2)

	// Convertir un string que representa un número a un float64
	weight := "70.5"
	weight2, _ := strconv.ParseFloat(weight, 64)

	fmt.Printf("Tipo: %T - Valor: %v", weight2, weight2)

}
