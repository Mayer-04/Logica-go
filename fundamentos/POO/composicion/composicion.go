package main

import "fmt"

//* Go no tiene herencia, utiliza la composición en su lugar.
// La composición se logra a través de la `incrustación` de tipos.

type Engine struct {
	Horsepower int
}

func (e Engine) Start() {
	fmt.Println("Engine started")
}

// La estructura Car incrusta el struct "Engine" como un campo con los métodos y campos definidos en Engine.
type Car struct {
	Engine
	Model string
}

func main() {
	myCar := Car{
		Engine: Engine{Horsepower: 200},
		Model:  "Toyota",
	}

	fmt.Println("Model:", myCar.Model)
	fmt.Println("Horsepower:", myCar.Horsepower)
	myCar.Start()
}
