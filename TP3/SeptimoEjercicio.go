package main

import "fmt"

/*
	Invertir: Leer 10 números ingresados por el usuario e imprimirlos en orden inverso
	Estructura: Pila (por el principio LIFO)
*/

func main() {
	// Crear una pila como un slice de enteros
	var pila []int

	fmt.Println("Ingrese 10 números:")

	// Leer 10 números y apilarlos (push)
	for i := 1; i <= 10; i++ {
		var num int
		fmt.Printf("Número %d: ", i)
		fmt.Scan(&num)

		// Apilar el número al final del slice
		pila = append(pila, num)
	}

	fmt.Println("\nNúmeros en orden inverso:")

	// Desapilar (pop) e imprimir
	for i := 0; len(pila) > 0; i++ {
		// Obtener el último elemento (top de la pila)
		top := pila[len(pila)-1]

		// Imprimir el valor
		fmt.Println(top)

		// Sacar el último elemento
		pila = pila[:len(pila)-1]
	}
}
