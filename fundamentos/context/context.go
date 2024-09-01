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
- Usa plazos de tiempo (deadlines) o tiempos de espera (timeouts) para evitar operaciones
que duren indefinidamente.
- Evita almacenar datos grandes o sensibles en contextos. Usa el contexto solo para
pasar valores que son necesarios para la operación.

Los contextos en Go son útiles para:
1. Pasar valores entre funciones.
2. Manejar cancelaciones en operaciones concurrentes.
3. Establecer plazos de tiempo (deadlines) o tiempos de espera (timeouts).
*/

// Definir un tipo de clave específico para evitar colisiones.
type requestIDKeyType string

const requestIDKey requestIDKeyType = "requestID"

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
	defer cancel() // Asegurar que los recursos de `ctxCancel` se liberen.

	// Crear un contexto derivado de `ctxCancel` con un tiempo de espera (timeout).
	// El contexto `ctxTimeout` se cancelará automáticamente después de 1 segundo
	// o si se llama a `cancel()`.
	ctxTimeout, cancelTimeout := context.WithTimeout(ctxCancel, time.Second)
	defer cancelTimeout() // Asegurar que los recursos de `ctxTimeout` también se liberen.

	// Simular una operación que podría respetar el contexto de tiempo de espera.
	select {
	case <-time.After(2 * time.Second): // Simular una operación que dura 2 segundos.
		fmt.Println("Operación completada.")
	case <-ctxTimeout.Done():
		// Se ejecuta si el contexto se cancela o se agota el tiempo.
		fmt.Println("Operación cancelada o tiempo agotado:", ctxTimeout.Err())
	}

	// Crear un contexto con un plazo de tiempo (deadline) específico.
	// Este contexto se cancelará cuando se alcance la fecha límite (deadline) o si se llama a `cancel()`.
	deadline := time.Now().Add(500 * time.Millisecond)
	ctxDeadline, cancelDeadline := context.WithDeadline(ctx, deadline)
	defer cancelDeadline()

	// Pasar valores a través de contextos:
	// Los valores pueden ser almacenados en un contexto y ser accesibles en funciones descendientes.
	// Esto es útil para pasar información como IDs de usuario, tokens de autenticación, etc.
	// Se recomienda que las claves no deben pasarse como strings.
	ctxValue := context.WithValue(ctxDeadline, requestIDKey, "12345")

	// Simular operaciones concurrentes que usan los contextos anteriores.
	go processRequest(ctxTimeout)
	go processRequest(ctxValue)

	// Esperar para permitir que las goroutines terminen.
	time.Sleep(2 * time.Second)
}

// processRequest Simula una función que procesa una solicitud respetando el contexto proporcionado.
func processRequest(ctx context.Context) {
	select {
	case <-time.After(700 * time.Millisecond): // Simula una operación que dura 700ms.
		if requestID := ctx.Value("requestID"); requestID != nil {
			fmt.Println("Operación completada con RequestID:", requestID)
			return
		}
		fmt.Println("Operación completada.")
	case <-ctx.Done():
		// Se ejecuta si el contexto se cancela o se agota el tiempo.
		fmt.Println("Operación cancelada:", ctx.Err())
	}
}

// func processRequest(ctx context.Context) {
// 	select {
// 	case <-time.After(700 * time.Millisecond):
// 		if requestID, ok := ctx.Value(requestIDKey).(string); ok {
// 			fmt.Printf("Operation completed with RequestID: %s\n", requestID)
// 			return
// 		}
// 		fmt.Println("Operation completed.")
// 	case <-ctx.Done():
// 		fmt.Println("Operation canceled:", ctx.Err())
// 	}
// }
