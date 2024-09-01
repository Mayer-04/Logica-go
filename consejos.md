# Consejos y Buenas Prácticas en Go

## Punto y Coma

En Go, los puntos y comas (;) son opcionales y el compilador los inserta automáticamente durante la compilación.

**Algunas caracteristicas a tener en cuenta:**

- No se insertara un punto y coma justo despues de la palabra reservada `for`.
- Con los operadores de incremento `++` y decremento `--` hay que tener cuidado y tratarlos como una declaración y no como una expresión.

La razón por la cual el código anterior no es válido es que los compiladores lo verán como:

```go
func main() {
    number := 0;
    println(number++;); // Error: ++ no se puede usar como expresión
    println(number--;); // Error: -- no se puede usar como expresión
}
```

Para que sea correcto, los operadores de incremento y decremento deben estar en su propia línea como declaraciones:

```go
func main() {
    number := 0
    number++            // Incrementa 'number' en 1
    println(number)     // Ahora 'number' es 1
    number--            // Decrementa 'number' en 1
    println(number)     // Ahora 'number' es 0
}
```

## Nombres de Paquetes

Los nombres de paquetes en Go deben ser cortos, claros y en minúsculas, evitando convenciones como `snake_case` o `camelCase`. Nombres como `utils`, `common,` `helpers`, etc., deben evitarse ya que no indican claramente la funcionalidad del paquete.

# La Mejor Forma de Inicializar un Slice en Go

Inicializar correctamente un slice en Go es crucial para el rendimiento de tu programa. Una mala inicialización puede llevar a reasignaciones innecesarias de memoria y sobrecargar el recolector de basura (GC), lo que afecta negativamente la eficiencia.

En Go, la función `make()` es la herramienta clave para inicializar slices de manera eficiente, permitiéndote especificar tanto la longitud como la capacidad.

Si sabes cuántos elementos tendrá el slice desde el inicio, define su longitud. Esto previene reasignaciones de memoria durante su uso.

```go
// Inicialización con longitud conocida
slice := make([]int, 10) // Crea un slice con 10 elementos, longitud 10 y capacidad 10
```

Si planeas añadir más elementos al slice, es recomendable inicializarlo con una longitud inicial de 0 pero con una capacidad mayor. Esto evita reasignaciones al utilizar la función `append()`.

```go
// Inicialización con longitud 0 pero capacidad 10
slice := make([]int, 0, 10) // Longitud 0, capacidad 10

for i := 1; i <= 10; i++ {
    slice = append(slice, i)
}

fmt.Println(slice) // Output: [1 2 3 4 5 6 7 8 9 10]
```

En este ejemplo, el slice se inicializa con longitud 0 pero con capacidad 10, permitiendo añadir elementos sin reasignaciones hasta alcanzar la capacidad.

### Ventajas de Inicializar un Slice con Mayor Capacidad

Inicializar un slice con una capacidad mayor a su longitud inicial es beneficioso cuando planeamos añadir elementos. Esto evita múltiples reasignaciones de memoria y mejora el rendimiento.

```go
// Inicialización con capacidad mayor a la longitud inicial
slice := make([]int, 5, 10) // Longitud inicial de 5 y capacidad de 10
```

### Conclusión

- Si conocemos la longitud final del slice y no planeamos usar append para agregar elementos, inicializarlo con una longitud específica es más eficiente.
- Si el slice va a crecer dinámicamente, inicializarlo con una capacidad que prevea el crecimiento puede evitar múltiples reasignaciones de memoria, mejorando el rendimiento.

## Desempaquetar Elementos de un Slice

Cuando utilizas los tres puntos (...) después de un slice, estás desempaquetando los elementos del slice, es decir, estás expandiendo el slice en una lista de sus elementos individuales. Esto es útil en situaciones donde necesitas pasar cada elemento del slice como un argumento separado a una función.

```go
package main

import "fmt"

func main() {
    original := []int{1, 2, 3, 4, 5}

    // Crear un slice vacío
    emptySlice := []int{}

    // Usar append con los tres puntos para desempaquetar los elementos del slice
    copySlice := append(emptySlice, original...)

    fmt.Println("Original slice:", original) // [1 2 3 4 5]
    fmt.Println("Copied slice:", copySlice)  // [1 2 3 4 5]
}
```

## Manejo de Errores en Go con fmt.Errorf, el Verbo %w y errors.Join

Desde Go 1.13, puedes añadir contexto a los errores sin perder la información original usando `fmt.Errorf` con el verbo `%w`. En Go 1.20, se introdujo `errors.Join` para agrupar múltiples errores.

### Formato de los Mensajes de Error

Es recomendable escribir los mensajes de error en **minúsculas** para mantener la consistencia, especialmente cuando los mensajes de error se combinan o se muestran juntos. A diferencia de otros lenguajes de programación donde es común **capitalizar** la primera letra de los mensajes de error, en Go se prefiere que comiencen en **minúscula**, salvo que inicien con un _nombre propio_ o una _abreviatura_.

**Envolver un error:**

```go
fmt.Errorf("failed to do something: %w", err)
```

**Unir errores**:

```go
errors.Join(err1, err2)
```

**¿Por Qué Envolver Errores?**

Envolver errores proporciona más contexto sobre dónde y por qué ocurrió el error, lo que facilita la depuración. El verbo %w permite mantener el error original, útil para inspección con `errors.Is` y `errors.As`.

## Evitar el uso de panic() y más en entornos de producción

Aunque `panic` puede ser recuperado con `recover()`, esto no siempre es posible ni recomendable.

- `recover()` solo puede capturar pánicos en la misma goroutine donde se llama, el pánico de otra función no pueden ser recuperado por la función diferida en main(). Esto resulta en un crash del programa.
- Un pánico en una parte del sistema puede provocar fallos en otras partes, especialmente en entornos de microservicios o sistemas distribuidos, causando fallos en cascada.
- Debería limitarse a situaciones críticas, como durante la inicialización de un programa.
