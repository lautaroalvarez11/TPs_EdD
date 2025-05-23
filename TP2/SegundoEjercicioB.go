package main

import "fmt"

/*
	Suma de enteros: Permite sumar todos los números enteros comprendidos entre un parámetro
	de inicio y uno de fin
*/

func sumarEntre(inicio, fin int) int {
	// Caso base: Si son iguales no hay números entre sí
	if inicio == fin {
		return inicio
	}

	return inicio + sumarEntre(inicio+1, fin) // Llamada recursiva
}

func sumarEntreComprendidos(inicio int, fin int) int {
	return sumarEntre(inicio+1, fin-1)
}

func main() {
	// Declarar dos variables para los parámetros
	var inicio, fin int

	// Pedir al usuario que ingrese los valores
	fmt.Print("Ingrese el número de inicio: ")
	fmt.Scan(&inicio)

	fmt.Print("Ingrese el número de fin: ")
	fmt.Scan(&fin)

	// Variable para almacenar la suma total
	suma := sumarEntreComprendidos(inicio, fin)

	// Mostrar el resultado
	fmt.Println("La suma de los números entre los parámetros es: ", suma)
}
