package main

import (
	"fmt"
	"sync"
)

func main() {
	waitGroup()
}

func waitGroup() {
	// Crea un WaitGroup para sincronizar las gorutinas
	var wg sync.WaitGroup

	// Incrementa el contador del WaitGroup en 1, indicando que vamos a esperar a que una goroutine termine
	wg.Add(1)

	go func() {
		// Decrementa el contador del WaitGroup en 1 cuando esta goroutine termine
		defer wg.Done()

		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
	}()

	// Bloquea el hilo principal hasta que el contador del WaitGroup sea 0,
	// es decir, hasta que todas las goroutines que hemos añadido con Add() hayan terminado
	wg.Wait()

	// Este mensaje se imprime una vez que todas las goroutines han completado su ejecución
	fmt.Println("Done!")
}

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	waitGroup()
// }

// // waitGroup demuestra el uso de sync.WaitGroup para sincronizar gorutinas

// // waitGroup inicializa un WaitGroup y lanza una cantidad especificada de gorutinas
// // Cada gorutina imprime su índice antes de marcarse como completada
// // El WaitGroup espera a que todas las gorutinas terminen antes de imprimir "Done!"
// func waitGroup() {
// 	// Inicializa un WaitGroup
// 	var wg sync.WaitGroup

// 	// Número de gorutinas a lanzar
// 	const count = 5

// 	// Añade el número de gorutinas al WaitGroup
// 	wg.Add(count)

// 	// Lanza las gorutinas
// 	for i := 0; i < count; i++ {
// 		// Lanza una gorutina y pasa su índice como argumento
// 		go func(i int) {
// 			// Defer la llamada a Done(), marcando la gorutina como completada
// 			defer wg.Done()
// 			// Imprime el índice de la gorutina
// 			fmt.Println(i)
// 		}(i)
// 	}

// 	// Espera a que todas las gorutinas terminen
// 	wg.Wait()

// 	// Imprime "Done!" después de que todas las gorutinas hayan terminado
// 	fmt.Println("Done!")
// }
