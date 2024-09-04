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

Desde **Go 1.13**, puedes añadir contexto a los errores sin perder la información original usando `fmt.Errorf` con el verbo `%w`. En **Go 1.20**, se introdujo `errors.Join` para agrupar múltiples errores.

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

- `recover()` solo puede capturar pánicos en la misma goroutine donde se llama, el pánico de otra función no pueden ser recuperado por la función diferida en _main()_. Esto resulta en un crash del programa.
- Un pánico en una parte del sistema puede provocar fallos en otras partes, especialmente en entornos de microservicios o sistemas distribuidos, causando fallos en cascada.
- Debería limitarse a situaciones críticas, como durante la inicialización de un programa.

## Diferencias entre new(T) y &T{}

**Inlining (inlinear):** Insertar el `código` de una función dentro de otra función para evitar la llamada a la función hija, lo que puede mejorar el rendimiento al reducir la sobrecarga de la llamada a la función.

Cuando Go compila el código, intenta inlinear las funciones para mejorar el rendimiento. El uso de `new(T)` en lugar de `&T{}` tiene un costo de `inlining` más bajo porque new(T) es una operación más simple y directa, lo que facilita que Go realice la optimización.

- Tanto `new(T)` como `&T{}` crean un nuevo valor de `tipo T` y devuelven un puntero a ese valor.
- `new(T)` inicializa el valor con el valor cero del `tipo T`.
- `&T{}` permite la inicialización de los campos del `tipo T` en el mismo paso.

## Construir un Binario más Óptimo en Go

Al construir un binario en Go, puedes utilizar los flags `-s` y `-w` para reducir su tamaño. Estos flags son opciones para el compilador de Go que eliminan información no esencial para la ejecución del binario, como la información de depuración.

**Depuración:** Es el proceso de **identificar y corregir errores** o comportamientos inesperados en un programa.

### Flags para Reducir el Tamaño del Binario

- `-s:` Elimina la tabla de símbolos del binario.
- `-w:` Elimina la información de depuración.

- **Tabla de Símbolos:** Contiene información sobre las funciones y variables del código. Aunque no afecta la ejecución del programa, ayuda a los depuradores a entender el código y facilita la depuración.

- **Información de Depuración:** Incluye detalles como los nombres de las variables y las líneas de código. No afecta el funcionamiento del programa, pero facilita la detección de errores durante el desarrollo. Elimina la información sobre cómo se generan las llamadas y devoluciones de funciones (la "pila" del programa).

**Nota Importante:** La eliminación de la tabla de símbolos y la información de depuración no afecta la ejecución del binario. El programa funcionará igual, pero será más difícil de depurar en caso de errores.

### Cómo Usar los Flags en la Línea de Comandos

Para utilizar estos flags al construir tu binario, añádelos a la línea de comandos de `go build`. Aquí tienes ejemplos de cómo hacerlo:

Para construir un binario con el nombre del archivo de Go:

```bash
go build -ldflags="-s -w" example.go
```

Para construir un binario con un nombre específico:

```bash
go build -ldflags="-s -w" -o my-binary example.go
```

### Explicación de los Comandos

- `go build:` Es el comando utilizado en el lenguaje de programación Go para compilar el código fuente en un binario ejecutable.
- `-ldflags:` Permite pasar flags específicos al proceso de enlace (linking) del compilador. Esto se utiliza para optimizar el tamaño y el rendimiento del binario.
- `-o:` Es una opción que define el nombre del binario resultante.

### Desactivar CGO (Opcional)

`CGO` es una característica de Go que permite a tu programa interactuar con código escrito en `C` u otros lenguajes que pueden compilarse en un formato compatible con `C`.

Si no necesitas interoperabilidad con código `C`, puedes desactivar _CGO_ utilizando la variable de entorno `CGO_ENABLED=0`. Esto puede hacer que el binario generado sea más pequeño y más fácil de portar entre diferentes sistemas.

```bash
CGO_ENABLED=0 go build -ldflags="-s -w" -o my-binary example.go
```

#### Beneficios de Desactivar CGO:

- **Binario más pequeño:** El compilador de Go no incluye dependencias ni código C en el binario resultante. Esto generalmente da como resultado un binario más ligero, ya que todo el código está escrito en Go puro.
- **Compilación más rápida:** Al no tener que enlazar con bibliotecas C ni ejecutar un compilador C, el proceso de compilación suele ser más rápido.
- Especialmente útil en entornos controlados como contenedores `Docker`, donde la simplicidad y la portabilidad son clave.

#### Consideraciones al Desactivar CGO:

- **Dependencias en C:** Si tu proyecto depende de bibliotecas escritas en C, no podrás usarlas si CGO está desactivado.
- **Funcionalidad del sistema operativo:** Algunas características específicas del sistema operativo, especialmente las que son de bajo nivel, pueden no estar disponibles directamente en Go sin CGO.

#### ¿Cuándo Desactivar CGO?

**Desactivar CGO es recomendable cuando:**

- Estás desarrollando aplicaciones que no necesitan interactuar con código C, como `microservicios` o `aplicaciones web`, donde todo el trabajo se realiza en Go.
- Necesitas crear un binario que funcione en múltiples plataformas sin cambios, evitando problemas de compatibilidad.
