package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

/*
* Sincronización de Goroutines con WaitGroup en Go
El paquete "sync" proporciona primitivas para sincronizar goroutines, entre ellas el `WaitGroup`.

El objetivo de usar un `WaitGroup` es asegurarse de que todas las goroutines terminen antes de que el programa principal
continúe, pero no se controla el orden en que se ejecutan las goroutines.

- `sync.WaitGroup`: Estructura que permite esperar a que un grupo de goroutines termine su ejecución.
- `Add(delta int)`: Incrementa el contador del WaitGroup por el valor de `delta`.
- `Done()`: Decrementa el contador del WaitGroup en 1. Debe ser llamado cuando una goroutine termina.
- `Wait()`: Bloquea la ejecución hasta que el contador del WaitGroup llega a 0, es decir,
hasta que todas las goroutines hayan llamado a `Done()`.
*/

func main() {

	waitGroup()
	fetchAllAPIs()
}

// waitGroup inicializa un WaitGroup, lanza múltiples goroutines y espera a que todas completen su ejecución.
func waitGroup() {
	// Crea una instancia de WaitGroup para manejar la sincronización de las gorutinas.
	var wg sync.WaitGroup

	// Definimos una constante para indicar el número de gorutinas que vamos a lanzar.
	const numGoroutines = 5
	// Incrementa el contador del WaitGroup en 5, indicando que vamos a esperar a que 5 gorutinas terminen.
	// NOTA: Llama al método `Add()` antes de que se llame al método `Done()`, no utilizarlo dentro de las gorutinas.
	wg.Add(numGoroutines)

	// Lanza 5 gorutinas que imprimen su índice antes de finalizar.
	for i := 1; i <= numGoroutines; i++ {
		go func(i int) {
			// Decrementa el contador del WaitGroup cuando la goroutine finaliza.
			// Cuando usas defer, te aseguras de que wg.Done() se ejecute justo antes de que la goroutine termine.
			// Nos ayuda a no preocuparnos en donde colocar el wg.Done() en el código.
			defer wg.Done()

			// Imprime el índice de la goroutine actual.
			fmt.Println("Goroutine:", i)
		}(i)
	}

	// Bloquea el hilo principal (main) hasta que el contador del WaitGroup sea 0,
	// es decir, hasta que todas las goroutines que hemos añadido con `Add()` hayan terminado.
	wg.Wait()

	// Imprime un mensaje indicando que todas las goroutines han finalizado.
	fmt.Println("Todas las goroutines se han completado!")
}

//* Consumiendo múltiples APIs con WaitGroup:

// URLs de las 3 APIs que queremos solicitar.
var URLS = [3]string{
	"https://jsonplaceholder.typicode.com/posts/1",
	"https://rickandmortyapi.com/api/character",
	"https://jsonplaceholder.typicode.com/users",
}

func fetchAllAPIs() {
	// Inicializa un WaitGroup para esperar hasta que todas las solicitudes a las APIs hayan finalizado.
	var wg sync.WaitGroup

	// Incrementa el contador del WaitGroup con la longitud de la lista de APIs.
	wg.Add(len(URLS))

	// Lanza una gorutina por cada API en el slice.
	for _, url := range URLS {
		go func(url string) {
			// Decrementa el contador del WaitGroup cuando cada gorutina finaliza.
			defer wg.Done()
			// Llama a la función requestAPI para realizar la solicitud HTTP.
			data, err := requestAPI(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			// Imprime los datos recibidos de la APIs como una cadena de texto.
			fmt.Println(string(data))
		}(url) // Pasamos la URL actual de la API a la función anonima (gorutina).
	}

	// Espera a que todas las gorutinas hayan completado su ejecución antes de continuar.
	wg.Wait()
}

func requestAPI(url string) ([]byte, error) {
	// Realiza una solicitud HTTP a la URL proporcionada y devuelve el cuerpo de la respuesta como un slice de bytes.
	// Si ocurre un error, devuelve un error detallado.
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error al realizar la solicitud: %w", err)
	}

	// Asegura que el cuerpo de la respuesta HTTP se cierre al finalizar la función, evitando fugas de recursos.
	defer resp.Body.Close()

	// Verifica que el estado de la respuesta HTTP sea 200 OK.
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("el estado de la respuesta es: %v", resp.StatusCode)
	}

	// Lee el cuerpo de la respuesta HTTP y devuelve los datos como un slice de bytes.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer los datos de la API: %w", err)
	}

	// Devuelve los datos leídos como un slice de bytes y un error nil, indicando que la operación fue exitosa.
	return body, nil
}
