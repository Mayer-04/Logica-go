package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

type Person struct {
	Name string
	Age  int
}

// La función antes de su nombre tiene un parámetro adicinal, este se conoce como "receptor"
// El receptor se coloca entre paréntesis, p se refiere al receptor que es una instancia a la struct Person
func (p Person) saludar() string {
	return "Hola" + " " + p.Name
}

//* Punteros en métodos
// No necesita un puntero en la variable de receptor cuando el método simplemente accede a la información del receptor.
func (p *Product) modificar() {
	p.Name = "iPhone"
}

func main() {

	// Antes de crear un método tienes que crear una instancia a la struct
	mayer := Person{Name: "Mayer", Age: 24}
	fmt.Println(mayer.saludar())

	celular := Product{Name: "Xiaomi 12", Price: 2.340}
	celular.modificar()
	fmt.Println(celular)
}
