package main

import (
	"context"
	"fmt"
	"time"
)

// TODO: Pasar valores de una función a otra - De una goroutine a otra
// TODO: WithDeadline: Fecha limite

// * Context Propagation

// Un contexto puede derivar de un contexto padre
// TODO: Es buena practica cancelar los contextos para liberar los recursos del contexto

func main() {

	// Crear un contexto vacio que no puede ser cancelado, contexto PADRE
	ctx := context.Background()

	// Este contexto se deriva del contexto padre "ctx"
	ctxCancel, cancel := context.WithCancel(ctx)

	// Cuando la función main termine se ejecuta la función "cancel"
	defer cancel()

	ctxTimeout := context.WithTimeout(ctxCancel, time.Second)

	fmt.Println(ctxTimeout)

}
