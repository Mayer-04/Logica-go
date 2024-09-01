package main

import (
	"context"
	"fmt"
	"time"
)

/*
* Paquete context en Go
El paquete `context` en Go es utilizadopara manejar la propagación
de cancelaciones y plazos de tiempo (deadlines) en operaciones concurrentes.

* Context Propagation: Crear y propagar contextos.
- Los contextos en Go son útiles para pasar valores entre funciones y para manejar cancelaciones
en operaciones concurrentes, como en goroutines. Un contexto puede derivarse de otro contexto
para propagar cancelaciones o fechas límite (deadlines).
*/

func main() {
	// Crear un contexto de fondo (background) que es el contexto raíz (padre).
	// Este contexto no puede ser cancelado y no tiene un plazo de tiempo asociado.
	ctx := context.Background()

	// Crear un contexto derivado de `ctx` que se puede cancelar.
	// El contexto `ctxCancel` hereda el valor de `ctx` y puede ser cancelado
	// para liberar recursos cuando ya no se necesita.
	ctxCancel, cancel := context.WithCancel(ctx)

	// `cancel()` se llama cuando la función `main` termine, para liberar los recursos asociados con `ctxCancel`.
	// Es una buena práctica llamar a `cancel` para evitar fugas de recursos.
	defer cancel()

	// Crear un contexto derivado de `ctxCancel` con un tiempo de espera (timeout).
	// El contexto `ctxTimeout` se cancelará automáticamente después de 1 segundo
	// o si se llama a `cancel()`.
	ctxTimeout, cancelTimeout := context.WithTimeout(ctxCancel, time.Second)
	defer cancelTimeout() // Asegurar que los recursos de `ctxTimeout` también se liberen.

	// Simular una operación que podría respetar el contexto de tiempo de espera.
	select {
	case <-time.After(3 * time.Second): // Simular una operación que dura 3 segundos.
		fmt.Println("Operación completada.")
	case <-ctxTimeout.Done():
		// Se ejecuta si el contexto se cancela o se agota el tiempo.
		fmt.Println("Operación cancelada o tiempo agotado:", ctxTimeout.Err())
	}
}
