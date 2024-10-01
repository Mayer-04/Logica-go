package main

import "fmt"

/*
Escribe una función en Python que calcule los primeros n términos de la secuencia de Fibonacci.
La secuencia de Fibonacci se define de la siguiente manera:

Los dos primeros términos son 0 y 1.
A partir del tercer término, cada número es la suma de los dos anteriores.
Por ejemplo, la secuencia de Fibonacci para n = 10 es:

0, 1, 1, 2, 3, 5, 8, 13, 21, 34
*/

// func main() {

// 	calcularFibonacci(15)

// }

// func calcularFibonacci(n int) {
// 	digi1 := 0
// 	digi2 := 1

// 	for i := 0; i < n; i++ {
// 		fmt.Print(digi1, " ")
// 		result := digi1 + digi2
// 		digi1 = digi2
// 		digi2 = result
// 	}
// }

func fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}

	fib := make([]int, n)
	fib[0] = 0

	if n > 1 {
		fib[1] = 1
	}

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib
}

func main() {
	n := 10
	result := fibonacci(n)
	fmt.Println(result)
}
