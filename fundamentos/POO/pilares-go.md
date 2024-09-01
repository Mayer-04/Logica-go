# Pilares de la Programación Orientada a Objetos en Go

## Abstracción

En Go, la abstracción se logra mediante la definición de `tipos y estructuras (structs)`. Cuando creas una estructura, seleccionas las propiedades que representan un concepto en tu programa.

**Ejemplo:**

Si estamos modelando un coche, podrías abstraerlo en una estructura _Car_ con campos como _Brand_, _Model_, y _Year_:

```go
type Car struct {
    Brand string
    Model string
    Year  int
}
```

## Encapsulamiento

El encapsulamiento en Go se logra utilizando la visibilidad de los campos y métodos. Los campos o métodos que comienzan con una letra mayúscula son públicos (accesibles desde otros paquetes), mientras que aquellos que comienzan con una letra minúscula son privados (accesibles solo dentro del mismo paquete).

```go
type Car struct {
    Brand string // Público
    model string // Privado
}
```

## Herencia

Go no tiene herencia como en otros lenguajes orientados a objetos, pero puedes lograr algo similar mediante la `composición`.

## Composición

La composición es la forma principal en que Go maneja la reutilización de código y la construcción de objetos complejos. En lugar de que una clase herede de otra, una estructura en Go puede incluir otra estructura como un campo, conocido como `incrustar una estructura`.

```go
type Engine struct {
    HorsePower int
}

// La estructura 'Car' contendra también los campos y métodos de la estructura 'Engine'.
type Car struct {
    Brand  string
    Model  string
    Year   int
    Engine // Composición: Incrustación de la estructura Engine en la estructura 'Car'.
}
```

## Polimorfismo

En Go, el polimorfismo se logra utilizando `interfaces`. Una interfaz define un conjunto de métodos, pero no proporciona la implementación de esos métodos.

- Cualquier tipo que implemente esos métodos se dice que "satisface" la interfaz.
- Una interfaz es como un contrato que deben cumplir los tipos en Go para que sean implementados.
- A diferencia de otros lenguajes de programación, no es necesario declarar _explícitamente_ que un tipo implementa una interfaz; Go lo determina automáticamente.

```go
type Vehicle interface {
    Start() string
}

// NOTA: La estructura 'Car' como la estructura 'Bike' implementan el método 'Start'de la interfaz 'Vehicle'.
type Car struct {
    Brand string
}

// Implementa de manera implícita la interfaz Vehicle.
func (c Car) Start() string {
    return "Car is starting"
}

type Bike struct {
    Brand string
}

// Implementa de manera implícita la interfaz Vehicle.
func (b Bike) Start() string {
    return "Bike is starting"
}

// StartVehicle puede trabajar con cualquier tipo que cumpla la interfaz 'Vehicle'.
func StartVehicle(v Vehicle) {
    fmt.Println(v.Start())
}

// Llamada al polimorfismo
car := Car{Brand: "Toyota"}
bike := Bike{Brand: "Yamaha"}

StartVehicle(car)
StartVehicle(bike)
```
