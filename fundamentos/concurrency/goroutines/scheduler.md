# Scheduler del Runtime de Go - Planificador en tiempo de ejecución de Go

En Go, las goroutines son unidades ligeras de ejecución gestionadas eficientemente por el `scheduler del runtime`. Aquí te explico cómo funciona:

## Gestión de Goroutines

El scheduler de Go maneja un montón de goroutines _(pequeños hilos ligeros)_ que se ejecutan en diferentes momentos. En lugar de asignar cada goroutine a un hilo físico, el scheduler organiza muchas goroutines para que usen unos pocos hilos del sistema operativo de manera eficiente.

El scheduler utiliza un algoritmo basado en el `modelo M`.

### Modelo M

Es una forma de organizar y gestionar cómo las goroutines (pequeños hilos ligeros) se ejecutan en los hilos del sistema operativo.

- M representa el número de goroutines que tienes.
- N representa el número de hilos del sistema operativo que tu computadora puede manejar a la vez.

En lugar de que cada goroutine tenga su propio hilo (lo que sería ineficiente y consumiría muchos recursos), el scheduler de Go toma muchas goroutines **(M)** y las distribuye en un número más pequeño de hilos del sistema operativo **(N)**. Esto permite que se aprovechen mejor los recursos de tu computadora sin sobrecargarla.

**Ejemplo de la vida real:**

Imagina que tienes 100 tareas **(M = 100)** pero solo 10 trabajadores **(N = 10)**. En lugar de asignar un trabajador a cada tarea, haces que esos 10 trabajadores se turnen para realizar las 100 tareas. Así no tienes que contratar 100 trabajadores para hacer el mismo trabajo.

## Colas de Goroutines

Cuando se crea una goroutine, se coloca en una especie de lista de espera llamada cola. El scheduler elige goroutines de esta lista para ejecutarlas en los hilos disponibles. Hay diferentes colas para goroutines que están listas para ejecutarse y para aquellas que están esperando por algo.

## Asignación a Hilos

Las goroutines se asignan a los hilos del sistema operativo disponibles mediante un conjunto de hilos gestionado por el runtime, conocido como "P". Estos hilos están asociados con los hilos del sistema operativo y permiten la ejecución concurrente de múltiples goroutines sin necesidad de asignar un hilo exclusivo para cada una.

## Operaciones Bloqueantes

Si una goroutine realiza una operación que la bloquea (como una espera en red o en un recurso), el scheduler puede pausar esa goroutine y moverla a una cola de espera. El scheduler entonces pone en espera esta goroutine y permite que otras goroutines continúen ejecutándose en el mismo hilo, evitando que el sistema se quede parado.

## Gestión de Colas

Si hay muchas goroutines esperando en la cola, el scheduler organiza las goroutines en varias colas para distribuir el trabajo entre los hilos. Esto ayuda a que el trabajo se haga de manera más equilibrada y rápida, evitando que algunas goroutines se queden esperando demasiado tiempo.
