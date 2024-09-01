package main

import "fmt"

/*
* Polimorfismo en Go
El polimorfismo en Go se logra a través de las 'interfaces'.
Esto permite que diferentes tipos compartan el mismo conjunto de métodos, lo que facilita que una función trabaje con
cualquier tipo que implemente la interfaz.

Características del polimorfismo en Go:
- Si un tipo implementa todos los métodos de una interfaz, se dice que ese tipo satisface la interfaz.
- A diferencia de otros lenguajes, no es necesario declarar explícitamente que un tipo implementa una interfaz;
Go lo determina automáticamente.
- Las interfaces permiten que diferentes tipos puedan ser tratados de la misma manera,
lo que facilita la reutilización del código.
*/

// Definimos una interfaz 'Vehicle' que tiene un método 'Start'.
// Cualquier tipo que implemente este método, automáticamente satisface la interfaz 'Vehicle'.
type Vehicle interface {
	Start() string
}

// Definimos una estructura 'Car' vacía que representa un tipo de vehículo.
type Car struct{}

// Implementamos el método 'Start' para la estructura 'Car'.
// Esto hace que 'Car' satisfaga la interfaz 'Vehicle'.
func (c Car) Start() string {
	return "Car is starting with a roar!"
}

// Definimos una estructura 'Motorcycle' vacía que representa otro tipo de vehículo.
type Motorcycle struct{}

// Implementamos el método 'Start' para la estructura 'Motorcycle'.
// Esto hace que 'Motorcycle' también satisfaga la interfaz 'Vehicle'.
func (m Motorcycle) Start() string {
	return "Motorcycle is starting with a vroom!"
}

// Definimos una estructura 'Bicycle' vacía que representa otro tipo de vehículo.
type Bicycle struct{}

// Implementamos el método 'Start' para la estructura 'Bicycle'.
// Esto hace que 'Bicycle' satisfaga la interfaz 'Vehicle'.
func (b Bicycle) Start() string {
	return "Bicycle is starting silently."
}

// La función 'StartVehicle' toma como parámetro cualquier tipo que implemente la interfaz 'Vehicle'.
// Imprime el resultado del método 'Start' del vehículo que recibe, demostrando el polimorfismo.
func StartVehicle(v Vehicle) {
	fmt.Println(v.Start())
}

func main() {

	// Creación de instancias de diferentes tipos que satisfacen la interfaz 'Vehicle'.
	// Esto demuestra cómo 'StartVehicle' puede trabajar con diferentes tipos de vehículos.
	// Todos estos tipos implementan la misma interfaz 'Vehicle'.
	car := Car{}
	StartVehicle(car) // Output: Car is starting with a roar!

	motorcycle := Motorcycle{}
	StartVehicle(motorcycle) // Output: Motorcycle is starting with a vroom!

	bicycle := Bicycle{}
	StartVehicle(bicycle) // Output: Bicycle is starting silently.
}
