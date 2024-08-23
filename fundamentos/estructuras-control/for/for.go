package main

import "fmt"

/*
* Bucle For
- El bucle `for` se utiliza para ejecutar un conjunto de instrucciones mientras una condición sea verdadera.
- En Go solo existe un tipo de bucle "for", el cual puede ser utilizado de diferentes maneras.
- Sintaxis: for inicialización; condición; actualización { bloque de código }
- El bucle `for` se puede usar para iterar sobre colecciones como arrays, strings, slices, maps y canales.
*/

func main() {

	// Bucle "for" clásico
	// Sintaxis: inicialización; condición; actualización
	fmt.Println("Bucle clásico:")
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	//* Bucle tipo "while"
	// Go no tiene un bucle "while" explícito, pero se puede emular con "for"
	i := 0
	for i <= 3 {
		fmt.Println("Bucle tipo while", i)
		i++
	}

	//* Bucle "for" con "range"
	// Se utiliza para iterar sobre colecciones como arrays, strings, slices, maps y canales.
	// Si solo necesitas el índice, puedes omitir el segundo valor.
	// Si solo necesitas el valor, puedes usar el identificador blank _ para omitir el índice. (_,v)
	numeros := [4]int{1, 2, 3, 4}
	// La variable i representa el índice y la variable v representa el valor
	for i, v := range numeros {
		fmt.Printf("Índice: %d, Valor: %d\n", i, v)
	}

	//* Bucle for con range de enteros
	// Desde la versión 1.22 de Go, se puede utilizar "range" con un número entero como argumento
	// Itera desde 0 hasta n-1
	for i := range 3 {
		fmt.Println(i) // Output: 0, 1, 2
	}

	//* Bucle for con "break"
	// Se detiene el bucle cuando i es igual a 50
	for i := 1; i <= 100; i++ {
		// Cuando i sea igual a 50, rompe el bucle
		if i == 50 {
			break
		}
		// Imprime los números del 1 al 49
		fmt.Println(i)
	}

	//* Bucle for con "continue"
	// Salta la iteración cuando i es igual a 4
	for i := 0; i <= 5; i++ {
		// Cuando i sea igual a 4, imprime "Continuando..." y salta a la siguiente iteración
		if i == 4 {
			fmt.Println("Continuando...")
			continue
		}
		// Imprime el mensaje para los valores de i que no son 4
		fmt.Printf("Soy el número %d\n", i)
	}

	//* Iterar sobre un string
	// Los caracteres son runas (runes), se imprime su valor Unicode.
	// Se debe hacer un casting a string para imprimir cada caracter como una cadena.
	nombre := "Mayer"
	for _, value := range nombre {
		fmt.Println(value)
	}

	//* Bucle infinito
	// Un bucle sin condiciones de salida explícitas se convierte en un bucle infinito
	// Se rompe el bucle con 'break' cuando count sea igual a 3
	count := 0
	for {
		fmt.Println("Bucle infinito, cuenta:", count)
		// Cuando count sea igual a 3, rompe el bucle
		if count == 3 {
			break // Salida del bucle infinito
		}
		count++
	}

	// Matrices
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	//* Bucles anidados para trabajar con una matriz (slice de slices)
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf("Elemento en [%d][%d] es %d\n", i, j, matrix[i][j])
		}
	}

	//* Bucle con más de una variable
	// La variable i y j se inicializan en 0 y 10 respectivamente
	// Su condición es que i sea menor que j
	// Su actualización es: i aumenta en 1 y j disminuye en 1
	// i imprime los números del 0 al 4 y j los imprime del 10 al 6
	for i, j := 0, 10; i < j; i, j = i+1, j-1 {
		fmt.Printf("i: %d, j: %d\n", i, j)
	}

	//* Simular un bucle "do while" en Go
	// Se asegura que el código se ejecute al menos una vez
	var numero int

	for {
		// Código que se ejecuta al menos una vez
		fmt.Print("Introduce un número positivo: ")
		// fmt.Scan -> Permite leer la entrada del usuario
		fmt.Scan(&numero)

		// Condición para romper el bucle en caso de que el número sea negativo con `break`
		if numero > 0 {
			break
		}
		fmt.Println("El número no es positivo. Inténtalo de nuevo.")
	}

	fmt.Printf("Número válido introducido: %d\n", numero)
}
