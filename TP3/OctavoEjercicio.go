package main

import (
	"fmt"
	"strings"
)

/*
	Palíndromo: Leer una secuencia de caracteres e informar si la misma es palíndromo.
	TAD: Pila
	Algoritmo:
	- Leer la cadena
	- Apilar todos los caracteres en una pila
	- Recorrer la cadena original, y a la vez desapilar elementos de la pila para compararlos:
		- Si todos coinciden → es un palíndromo.
		- Si alguno difiere → no es un palíndromo.
*/

func main() {
	var input string

	// Leer la cadena del usuario
	fmt.Print("Ingrese una cadena: ")
	fmt.Scanln(&input)

	// Convertir a minúsculas y eliminar espacios
	cadena := strings.ToLower(strings.ReplaceAll(input, " ", ""))

	var pila []rune
	for _, char := range cadena {
		pila = append(pila, char) // Push
	}

	// Comparar cada carácter con lo que se va desapilando
	esPalindromo := true
	for _, char := range cadena {
		top := pila[len(pila)-1]  // mirar el tope
		pila = pila[:len(pila)-1] // pop

		if char != top {
			esPalindromo = false
			break
		}
	}

	// Mostrar resultado
	if esPalindromo {
		fmt.Println("La cadena es un palíndromo.")
	} else {
		fmt.Println("La cadena no es un palíndromo.")
	}
}
