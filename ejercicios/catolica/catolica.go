package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	number := 45

	result := multiplesOperations(number)

	fmt.Println(result)
}

func multiplesOperations(number int) map[string]any {

	convertString := strconv.Itoa(number)

	fmt.Println(convertString)

	// TODO: Convertimos el byte a string
	firstNumber := string(convertString[0])
	secondNumber := string(convertString[1])

	fmt.Println(firstNumber, secondNumber)

	var newStr = ""
	for i := len(convertString) - 1; i >= 0; i-- {
		newStr += string(convertString[i])
	}

	fmt.Println(newStr)

	// TODO: Convertimos el string a int
	a, _ := strconv.Atoi(firstNumber)
	b, _ := strconv.Atoi(secondNumber)

	sum := a + b
	substraction := a - b
	multiplication := a * b
	division := float64(a) / float64(b)
	pow := int(math.Pow(float64(a), float64(b)))
	module := a % b

	return map[string]any{
		"reverse":        newStr,
		"sum":            sum,
		"substraction":   substraction,
		"multiplication": multiplication,
		"division":       division,
		"pow":            pow,
		"module":         module,
	}

}
