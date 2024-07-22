package main

import "fmt"

// Definición del struct "Product"
type Product struct {
	Name  string
	Price float64
}

// Definición del struct "Person"
type Person struct {
	Name     string
	Lastname string
	Age      int
}

// * La función saludar es un método del struct Person.
// El parámetro entre paréntesis antes del nombre de la función es el "receptor"
// El receptor (p Person) hace referencia a una instancia de la struct Person
// Los métodos con receptores de tipo valor (no puntero) trabajan con una copia del valor del receptor
func (p Person) saludar() string {
	return "Hola" + " " + p.Name
}

// * Punteros en métodos
// La función modificar es un método del struct Product
// El receptor (p *Product) es un puntero a una instancia de Product
// Los métodos con receptores de tipo puntero pueden modificar el valor del receptor
// Usar un receptor de tipo puntero es eficiente para structs grandes y cuando se quiere modificar el receptor
func (p *Product) modificar() {
	p.Name = "iPhone"
}

func main() {

	// Antes de invocar métodos en una estructura, debes crear una instancia de esa estructura
	// Esto se hace asignando valores a los campos del struct.
	mayer := Person{Name: "Mayer", Lastname: "Prada", Age: 24}
	fmt.Println(mayer.saludar()) // Output: Hola Mayer

	// Crear una instancia del struct Product e invocar su método modificar.
	celular := Product{Name: "Xiaomi 12", Price: 2.340} // Modifica el campo Name de la instancia celular
	celular.modificar()
	fmt.Println(celular) // Output: {iPhone 2340}

	// Nota: Cuando se llama a un método con receptor puntero, Go automáticamente toma la dirección de la variable
}
