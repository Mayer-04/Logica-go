package main

import (
	"errors"
	"fmt"
)

func main() {

	var numbers = make([]int, 0, 6)

	numbers = append(numbers, 1, 2, 3, 4, 5, 6)

	result, err := addEvenNumbers(numbers)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

}

// Suma de enteros pares: Escribe un programa que sume todos los números pares en un rango dado de enteros.
func addEvenNumbers(nums []int) (int, error) {

	if len(nums) == 0 {
		return 0, errors.New("no hay números")
	}

	result := 0

	for _, value := range nums {
		if value%2 == 0 {
			result += value
		}
	}

	return result, nil
}
