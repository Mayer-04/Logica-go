package main

import (
	"fmt"
	"time"
)

// * Las goroutines se crean con la palabra clave "go" antes del nombre de la función
func main() {

	go goroutine1()

	time.Sleep(time.Second * 1)
}

func goroutine1() {
	fmt.Println("Soy la primera goroutine")
}
