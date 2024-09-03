# Scheduler del Runtime de Go - Planificador en tiempo de ejecución de Go

En Go, las goroutines son unidades ligeras de ejecución que permiten la concurrencia de manera eficiente. En este artículo, te explico cómo el scheduler del runtime de Go gestiona estas goroutines.

## Gestión de Goroutines

El scheduler de Go maneja miles de goroutines que se ejecutan de manera concurrente. En lugar de asignar cada goroutine a un hilo físico del sistema operativo, el scheduler organiza muchas goroutines para que utilicen unos pocos hilos de manera eficiente

El scheduler utiliza el `Modelo GMP`.

### Modelo GMP: Goroutine, Machine, Processor

Go utiliza un modelo conocido como GMP para gestionar la ejecución de las goroutines:

- G (Goroutine): Es una tarea ligera que ejecuta código en Go.
- M (Machine): Representa un hilo del sistema operativo.
- P (Processor): Es una abstracción utilizada por el scheduler para manejar la ejecución de las goroutines. Cada P puede ejecutar una goroutine a la vez.

El número de M (hilos del sistema operativo) puede ser menor que el número de G (goroutines), lo que permite que muchas goroutines se ejecuten de manera eficiente utilizando pocos hilos.

**Ejemplo de la vida real:**

Imagina que tienes 100 tareas **(G = 100)** pero solo 10 trabajadores **(M = 10)**. En lugar de asignar un trabajador a cada tarea, esos 10 trabajadores se turnan para realizar las 100 tareas. El (P) actúa como un supervisor que organiza y distribuye las tareas a los trabajadores.

## Colas de Goroutines

Las goroutines se colocan en colas de espera antes de ejecutarse. Existen colas locales asociadas a cada (P) y una cola global. Las goroutines en las colas locales son las primeras en ser ejecutadas, y si un P no tiene más goroutines en su cola, puede tomar una de la cola global o *"robar"* una goroutine de otra cola local.

## Asignación a Hilos

Cada P tiene su propia cola de goroutines y se asocia con un M (hilo del sistema operativo) cuando ejecuta una goroutine. Si una goroutine realiza una operación que la bloquea, el M puede ser reasignado a otro P para continuar ejecutando otras goroutines, manteniendo la eficiencia.

## Operaciones Bloqueantes

Cuando una goroutine se bloquea (por ejemplo, esperando una respuesta de la red), el scheduler mueve esa goroutine a una cola de espera y reasigna el M que estaba ejecutándola a otra goroutine que esté lista para continuar.

## Gestión de Colas

El scheduler gestiona las colas de goroutines para asegurar que el trabajo se distribuye de manera equitativa entre los hilos disponibles, evitando que algunas goroutines queden esperando demasiado tiempo.
