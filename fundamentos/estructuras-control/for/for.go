package main

import "fmt"

func main() {

	//* En Go solo existe un tipo de bucle "for", el cual puede ser utilizado de diferentes maneras.

	//* Bucle for clásico
	// Sintaxis: inicialización; condición; actualización
	for i := 0; i <= 5; i++ {
		fmt.Println("Bucle clásico", i)
	}

	//* Bucle tipo while
	// Go no tiene un bucle while, pero se puede emular con for
	i := 0
	for i <= 3 {
		fmt.Println("Bucle tipo while", i)
		i++
	}

	//* Bucle for con range
	// Se utiliza para iterar sobre colecciones como arrays, strings, slices, maps y canales
	// Si solo necesitas el índice, puedes omitir el segundo valor.
	// Si solo necesitas el valor, puedes usar el identificador blank _ para omitir el índice.
	numeros := [4]int{1, 2, 3, 4}
	// La variable i representa el índice y la variable v representa el valor
	for i, v := range numeros {
		fmt.Printf("Índice: %d, Valor: %d\n", i, v)
	}

	//* Bucle for con range de enteros
	// Desde la versión 1.22 se puede utilizar range con un número entero como argumento
	// Itera desde 0 hasta n-1
	for i := range 3 {
		fmt.Println(i) // Output: 0, 1, 2
	}

	//* Bucle for con break
	for i := 1; i <= 100; i++ {
		// Cuando i sea igual a 50, rompe el bucle
		if i == 50 {
			break
		}
		// Imprime los números del 1 al 49
		fmt.Println(i)
	}

	//* Bucle for con continue
	for i := 0; i <= 5; i++ {
		// Cuando i sea igual a 4, imprime "Continuando..." y salta a la siguiente iteración
		if i == 4 {
			fmt.Println("Continuando...")
			continue
		}
		// Imprime el mensaje para los valores de i que no son 4
		message := fmt.Sprintf("Soy el número %d", i)
		fmt.Println(message)
	}

	//* Iterar un string
	// Se debe hacer un casting a string() porque los caracteres son runas, lo que nos retorna su valor Unicode
	var nombre string = "Mayer"
	for _, value := range nombre {
		fmt.Println(string(value))
	}

	//* Bucle infinito
	// Un bucle sin condiciones de salida explícitas se convierte en un bucle infinito
	count := 0
	for {
		fmt.Println("Bucle infinito, cuenta:", count)
		// Cuando count sea igual a 3, rompe el bucle
		if count == 3 {
			break // Salida del bucle infinito
		}
		count++
	}

	//* Bucles anidados para trabajar con una matriz (slice de slices)
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf("Elemento en [%d][%d] es %d\n", i, j, matrix[i][j])
		}
	}

}
