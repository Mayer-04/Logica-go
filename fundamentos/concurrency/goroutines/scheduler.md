# Scheduler del Runtime de Go - Planificador en tiempo de ejecución de Go

En Go, las goroutines son unidades ligeras de ejecución que permiten la concurrencia de manera eficiente. En este artículo, te explico cómo el scheduler del runtime de Go gestiona estas goroutines.

## Gestión de Goroutines

El scheduler de Go maneja miles de goroutines que se ejecutan de manera concurrente. En lugar de asignar cada goroutine a un hilo físico del sistema operativo, el scheduler organiza muchas goroutines para que utilicen unos pocos hilos de manera eficiente

El scheduler utiliza un modelo conocido como `GMP`.

### Modelo GMP: Goroutine, Machine, Processor

Go utiliza un modelo conocido como `GMP` para gestionar la ejecución de las goroutines:

- **G (Goroutine):** Es una tarea ligera que ejecuta código en Go.
- **M (Machine):** Representa un hilo del sistema operativo.
- **P (Processor):** Es un recurso dentro del scheduler de Go que organiza y ejecuta las goroutines. Piensa en `P` como un espacio de trabajo que permite a un **hilo (M)** ejecutar código Go. Cada `P` puede manejar una goroutine a la vez. Sin un `P`, un hilo no puede ejecutar código Go, incluso si está disponible. `P` se encarga de asignar y distribuir las tareas a los **hilos (M)** para que el trabajo se realice de manera eficiente.

_El número de M (hilos del sistema operativo) suele ser menor que el número de G (goroutines), lo que permite que muchas goroutines se ejecuten de manera eficiente utilizando pocos hilos._

**Cada goroutine (G) se ejecuta en un hilo del sistema operativo (M) que se asigna a un procesador lógico o CPU lógica (P).**

**Ejemplo de la vida real:**

Imagina que tienes 100 tareas **(G = 100)** pero solo 10 trabajadores **(M = 10)**. En lugar de asignar un trabajador a cada tarea, esos 10 trabajadores se turnan para realizar las 100 tareas. El **P** actúa como un supervisor que organiza y distribuye las tareas a los trabajadores.

**Ejemplo 2:**

Imagina que tienes una cocina con `100 recetas (las goroutines)` que preparar, pero solo `10 cocineros (los hilos)`. En lugar de asignar un cocinero a cada receta, esos 10 cocineros se turnan para trabajar en las 100 recetas. Cada cocinero tiene acceso a `estaciones de trabajo (los P)`, que son espacios donde pueden preparar las recetas. Las estaciones de trabajo organizan y distribuyen las recetas a los cocineros, asegurando que todo se haga de manera eficiente. Si un cocinero se bloquea esperando un ingrediente (una operación bloqueante), otro cocinero puede usar la misma estación de trabajo para continuar con otra receta.

## Colas de Goroutines

Las goroutines se colocan en colas de espera antes de ejecutarse. Existen colas locales asociadas a cada `P` y una cola global. Las goroutines en las colas locales son las primeras en ser ejecutadas, y si un `P` no tiene más goroutines en su cola, puede tomar una de la cola global o _"robar"_ una goroutine de otra cola local.

## Asignación a Hilos

Cada **procesador (P)** tiene su propia cola de goroutines y se asocia con un `hilo del sistema operativo` cuando ejecuta una goroutine. Si una goroutine realiza una operación que la bloquea, el `hilo` puede ser reasignado a otro `procesador` para continuar ejecutando otras goroutines, manteniendo la eficiencia.

## Operaciones Bloqueantes

Cuando una goroutine se bloquea (por ejemplo, esperando una respuesta de la red), el scheduler mueve esa goroutine a una cola de espera y reasigna el `hilo` que estaba ejecutándola a otra goroutine que esté lista para continuar.

## Gestión de Colas

El scheduler gestiona las colas de goroutines para asegurar que el trabajo se distribuye de manera equitativa entre los hilos disponibles, evitando que algunas goroutines queden esperando demasiado tiempo.
