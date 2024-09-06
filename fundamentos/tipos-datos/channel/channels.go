package main

import (
	"fmt"
	"time"
)

/*
* Channels: Canales en Go
Los canales en Go permiten la comunicación y sincronización entre goroutines, facilitando el envío y recepción de datos.

- Los canales deben crearse asignandoles memoria utilizando la función `make()`. Su tipo de dato es `chan`.
- El operador `<-` se usa para enviar datos a un canal (canal <- dato) y para recibir datos de un canal (dato <- canal).
- Los canales pueden bloquearse si no hay una goroutine disponible
para recibir los datos enviados o para enviar los datos requeridos.
- La función `close()` se utiliza para indicar que no se enviarán más datos a través de un canal.

* Existen canales con y sin búfer:
- Sin búfer: Los envíos y recepciones se bloquean hasta que la otra parte esté lista.
- Con búfer: Se pueden almacenar varios datos en el canal antes de que se bloquee.

* IMPORTANTE:
- Evitar el `deadlock`, que ocurre cuando todas las goroutines están bloqueadas, esperando unas de otras.
- Los canales unidireccionales se usan para limitar el uso de un canal a solo enviar o solo recibir datos,
lo cual puede ayudar a evitar errores.
*/

func main() {

	// Creando un canal de tipo string.
	ch := make(chan string)

	//* Enviar datos al canal desde una goroutine.
	go func() {
		time.Sleep(1 * time.Second) // Simulamos una operación que 1 segundo en ejecutarse.
		ch <- "Hello, World!"       // Envía el mensaje "Hello, World!" al canal `ch`.
	}()

	//* Recibir datos del canal.
	msg := <-ch // Espera a recibir un mensaje desde el canal y lo asigna a la variable `msg`.
	fmt.Println(msg)

	// * Canales unidireccionales.
	num := make(chan int)

	// Este canal solo puede recibir datos.
	go receive(num)

	// Este canal solo puede enviar datos.
	send(num)

	//* Creando un canal con búfer.
	// Este canal puede almacenar hasta 2 enteros antes de bloquearse.
	ch2 := make(chan int, 2)

	// Enviando datos al canal. No se bloquea porque tiene capacidad para almacenar dos valores.
	// Las operaciones de envío al canal solo se bloquearán cuando el búfer esté lleno.
	ch2 <- 4
	ch2 <- 2
	// ch2 <- 6 // Se bloquearía si se descomenta, ya que el búfer está lleno.

	// Recibiendo datos del canal.
	// Las operaciones de recepción solo se bloquearán cuando el búfer esté vacío.
	fmt.Println(<-ch2) // Imprime el primer valor almacenado en el canal (4).
	fmt.Println(<-ch2) // Imprime el segundo valor almacenado en el canal (2).

	// * Ejemplo de uso de canal cerrado.
	// Creando un canal de enteros.
	ch3 := make(chan int, 2)

	go func() {
		for i := 1; i <= 3; i++ {
			ch3 <- i // Enviando valores al canal.
		}
		close(ch3) // Cerrando el canal después de enviar todos los valores.
	}()

	// Recibiendo datos del canal hasta que esté cerrado.
	for val := range ch3 {
		fmt.Println(val) // Imprime los valores 1, 2, 3.
	}
}

// `send` envía un valor entero al canal.
// Esta función solo acepta un canal que envie datos.
func send(num chan<- int) {
	num <- 10
}

// `receive` recibe un valor entero del canal y lo imprime.
// Esta función solo acepta un canal que reciba datos.
func receive(num <-chan int) {
	fmt.Println(<-num)
}
