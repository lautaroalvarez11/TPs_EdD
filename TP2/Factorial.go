package main

import "fmt"

/*
	Factorial: Permite calcular el factorial de un número
*/

func factorial(n int) int {
	// Caso base: Si el número es 0, su factorial es 1
	if n == 0 {
		return 1
	}

	// Manejar números negativos (el factorial no está definido para negativos en enteros)
	if n < 0 {
		fmt.Println("Error: No se puede calcular el factorial de un número negativo")
		return -1 // Se devuelve un valor especial para indicar un error
	}

	// Caso recursivo: Multiplicar el número por el factorial del número anterior
	return n * factorial(n-1)
}

func main() {
	var num int

	fmt.Print("Ingrese un número para calcular el factorial: ")
	fmt.Scan(&num)

	// Variable para almacenar el resultado
	resultado := factorial(num)

	fmt.Println("El factorial de", num, "es:", resultado)
}
