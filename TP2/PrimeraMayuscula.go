package main

import (
	"fmt"
	"unicode"
)

/*
	Posición mayúscula: Dada una cadena de caracteres, devuelve la posición en la que se
	encuentra la primera letra mayúscula.
*/

func posicionMayuscula(cadena string, index int) int {
	// Caso base 1: Si la cadena está vacía o se recorrió toda la cadena, devolver -1 (no encontrada)
	if index >= len(cadena) {
		return -1
	}

	// Caso base 2: Si el carácter actual es mayúscula, devolver la posición
	if unicode.IsUpper(rune(cadena[index])) {
		return index
	}

	// Caso recursivo: Llamar a la función con el siguiente índice
	return posicionMayuscula(cadena, index+1)
}

func main() {
	var cadena string

	fmt.Print("Ingrese una cadena de caracteres: ")
	fmt.Scan(&cadena)

	posicion := posicionMayuscula(cadena, 0)

	fmt.Println("La posición de la primera mayúscula es:", posicion)
}
