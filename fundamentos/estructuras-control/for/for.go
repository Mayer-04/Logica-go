package main

import "fmt"

func main() {

	// TODO: En Go solo existe un tipo de bucle - FOR

	// Bucle for clasico
	// TODO: inicialización; condición; actualización
	for i := 0; i <= 5; i++ {
		fmt.Println("Bucle clasico", i)
	}

	// Bucle tipo while
	i := 0
	for i <= 3 {
		fmt.Println("while", i)
		i++
	}

	// Bluce for con range
	// TODO: Se utiliza para iterar colecciones como arrays, strings, slice, map, chan
	numeros := [4]int{1, 2, 3, 4}
	for i, v := range numeros {
		fmt.Printf("Indice: %d, Valores: %v\n", i, v)
	}

	// Bluce for con range de enteros
	// TODO: Desde la versión 1.22 se puede utilizar range con un número entero como argumento
	for i := range 3 {
		fmt.Println(i)
	}

	// Bucle for con break
	for i := 1; i <= 100; i++ {

		if i == 50 {
			break
		}
		fmt.Println(i)
	}

	// Bucle for con continue
	for i := 0; i <= 5; i++ {
		message := fmt.Sprintf("Soy el número %d", i)
		if i == 4 {
			fmt.Println(message)
			continue
		}
		fmt.Println(message)
	}

}
