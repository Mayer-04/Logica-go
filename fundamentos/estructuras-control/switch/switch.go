package main

import "fmt"

func main() {

	edad := 18

	// * Switch: una alternativa más limpia y eficiente a múltiples if-else
	// Evita la necesidad de múltiples if-else cuando se tiene una única variable a comparar.
	switch {
	case edad < 13:
		fmt.Println("Eres un niño")
	case edad < 18:
		fmt.Println("Eres un adolescente")
	default:
		fmt.Println("Eres un adulto")
	}

	// * Uso de etiquetas: evita la repetición de valores constantes
	// También se pueden utilizar etiquetas en switch para manejar diferentes casos con el mismo código.
	switch edad := edad; {
	case edad < 13:
		fmt.Println("Eres un niño")
	case edad < 18:
		fmt.Println("Eres un adolescente")
	default:
		fmt.Println("Eres un adulto")
	}

	var x interface{} = true

	// * Switch con tipo de variable
	// Se puede utilizar para el caso de que la variable sea de otro tipo
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
