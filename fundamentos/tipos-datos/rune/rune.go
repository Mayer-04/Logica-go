package main

import "fmt"

/*
* Rune: Runa
- Es un alias para el tipo int32.
- Puede representar caracteres Unicode que pueden ocupar más de un byte en la codificación UTF-8.
- El valor de una runa es su punto de código Unicode.
- Cuando se define un carácter como runa, Go usa su valor numérico Unicode.
*/

func main() {

	/*
		* Explicación de como convertir este valor Hexadecimal a Decimal:
		Como ejemplo tomares el valor de la runa 'a', si número en Unicode es: U+0061, su valor en decimal es 97.
		- Hexadecimal (base 16) utiliza los siguientes valores: 0-9, A-F.
		- Los dígitos A-F se representan como: A = 10, B = 11, C = 12, D = 13, E = 14, F = 15.
		- Cada dígito hexadecimal tiene 4 bits, por lo que el valor numérico de la runa es 16^3 + 16^2 + 16^1 + 16^0
		- El valor numérico de la runa es 16^3 + 16^2 + 16^1 + 16^0.

		1. 0 x 16^3 = 0
		2. 0 x 16^2 = 0
		3. 6 x 16^1 = 6 * 16 = 96
		4. 1 x 16^0 = 1 * 1 = 1
		5. 0 + 0 + 96 + 1 = 97

		La mejor fórmula de calcular el valor hexadecimal es:
		6 * 16^1 + 1 * 16^0 = 6 * 16 + 1 = 97
	*/

	// Definiendo una variable de tipo rune
	// Número en Unicode: U+0061
	runa := 'a'
	fmt.Println("runa:", runa) // Output: runa: 97

	runa2 := '🍎'
	fmt.Println("runa2:", runa2) // Output: runa2: 127822

}
