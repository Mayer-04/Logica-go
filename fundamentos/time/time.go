package main

import (
	"fmt"
	"time"
)

/*
* Paquete time en Go
El paquete "time" proporciona funcionalidades para manejar fechas, horas, duraciones y temporizadores.

- `time.Now()`: Obtiene la fecha y hora actual.
- `time.Add(d time.Duration)`: Suma una duración a una fecha y hora.
- `time.Sub(t time.Time)`: Calcula la diferencia entre dos instancias de tiempo.
- `time.Format(layout string)`: Formatea la fecha y hora a un string.
- `time.ParseDuration(s string)`: Convierte un string a un tipo `time.Duration`.
- `time.Sleep(d time.Duration)`: Pausa la ejecución durante una duración específica.
*/

func main() {

	// Obtener la fecha y hora actual.
	now := time.Now()
	fmt.Println("Fecha y hora actual:", now) // 2024-08-23 15:03:20.3843366 -0500 -05 m=+0.000000001

	// Sumar 2 horas a la fecha y hora actual.
	future := now.Add(2 * time.Hour)
	fmt.Println("Fecha y hora después de 2 horas:", future)

	// Calcular la diferencia de tiempo entre dos instancias.
	difference := future.Sub(now)
	fmt.Println("Diferencia entre future y now:", difference)

	// Formatear la fecha y hora actual en un string.
	formattedTime := now.Format("2006-01-02 15:04:05")
	fmt.Println("Fecha y hora formateada:", formattedTime)

	// Parsear un string a un tipo time.Duration.
	duration, err := time.ParseDuration("2h45m")
	if err != nil {
		panic(err) // Manejar el error adecuadamente en un contexto real.
	}
	fmt.Println("Duración parseada:", duration) // Duración parseada: 2h45m

	// Obtener el número de horas de una duración.
	hours := duration.Hours()
	fmt.Printf("Horas en la duración: %.2f\n", hours) // 2.75

	// Obtener el número de minutos de una duración.
	minutes := duration.Minutes()
	fmt.Printf("Minutos en la duración: %.2f\n", minutes) // 165.00

	// Obtener el número de segundos de una duración.
	seconds := duration.Seconds()
	fmt.Printf("Segundos en la duración: %.2f\n", seconds) // 9900.00

	// Pausar la ejecución del programa por 2 segundos.
	// Una duración negativa o cero hace que `Sleep` regrese inmediatamente.
	fmt.Println("Esperando 2 segundos...")
	time.Sleep(2 * time.Second)
	fmt.Println("Continuando la ejecución después de la pausa.")

	// Crear una fecha específica (1 de enero de 2024 a las 10:00 AM).
	specificDate := time.Date(2024, time.January, 1, 10, 0, 0, 0, time.UTC)
	fmt.Println("Fecha específica:", specificDate) // Fecha específica: 2024-01-01 10:00:00 +0000 UTC

	// Comparar dos fechas (after y before).
	if future.After(specificDate) {
		fmt.Println("El tiempo futuro es después de la fecha específica.") // Se imprime este println
	} else if future.Before(specificDate) {
		fmt.Println("El tiempo futuro es antes de la fecha específica.")
	} else {
		fmt.Println("El tiempo futuro es igual a la fecha específica.")
	}

	// Usar un temporizador para realizar una acción después de un tiempo.
	timer := time.NewTimer(3 * time.Second)
	fmt.Println("Temporizador de 3 segundos iniciado...")
	<-timer.C
	fmt.Println("Temporizador finalizado.")
}
