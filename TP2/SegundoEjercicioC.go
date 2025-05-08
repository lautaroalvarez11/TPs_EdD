package main

import (
	"fmt"
	"unicode"
)

/*
	Vocales y consonantes: Devuelve la cantidad de vocales y consonantes que contiene una
	cadena.
*/

func contarVocalesConsonantes(cadena string, index, vocales, consonantes int) (int, int) {
	// Caso base: Cuando se recorrió toda la cadena
	if index == len(cadena) {
		return vocales, consonantes
	}

	// Obtener el carácter actual y convertirlo a minúscula
	char := unicode.ToLower(rune(cadena[index]))

	// Verificar si es una letra
	if unicode.IsLetter(char) {
		// Verificar si es vocal
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			vocales++
		} else {
			consonantes++
		}
	}

	// Llamada recursiva para el siguiente carácter
	return contarVocalesConsonantes(cadena, index+1, vocales, consonantes)
}

func main() {
	var texto string

	fmt.Print("Ingrese una cadena de texto: ")
	fmt.Scanln(&texto)

	// Llamar a la función recursiva
	vocales, consonantes := contarVocalesConsonantes(texto, 0, 0, 0)

	// Mostrar los resultados
	fmt.Println("Cantidad de vocales:", vocales)
	fmt.Println("Cantidad de consonantes:", consonantes)
}
