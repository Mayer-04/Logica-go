package main

import "fmt"

func main() {

	var x interface{} = true
	switch x.(type) {
	case int:
		fmt.Println("Es un entero")
	case float64:
		fmt.Println("Es un flotante")
	case string:
		fmt.Println("Es una cadena")
	default:

		fmt.Println("Es de otro tipo")

	}

}
