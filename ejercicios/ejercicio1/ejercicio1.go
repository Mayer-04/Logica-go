package main

import (
	"errors"
	"fmt"
)

// Suma de enteros pares: Escribe un programa que sume todos los números pares en un rango dado de enteros.

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6}
	result, err := addEvenNumbers(numbers)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

}

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
