package main

import "fmt"

/*
	Mayor elemento: Dado un arreglo de enteros, devolver el mayor elemento.
*/

func mayorElemento(arr []int, index int, num int) int {
	// Caso base: Cuando se recorrió todo el arreglo
	if index == len(arr) {
		return num
	}

	// Caso recursivo: Comparar el elemento actual con el mayor
	if arr[index] > num {
		num = arr[index]
	}

	return mayorElemento(arr, index+1, num)
}

func main() {
	var n int

	fmt.Print("¿Cuántos elementos tendrá el arreglo?: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Ingrese los elementos del arreglo:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&arr[i])
	}

	if len(arr) == 0 {
		fmt.Println("El arreglo está vacío.")
		return
	}

	mayor := mayorElemento(arr, 0, arr[0])

	fmt.Println("El mayor elemento del arreglo es:", mayor)
}

/*
Forma Iterativa: O(n)

func main() {
	var n int

	fmt.Print("¿Cuántos elementos tendrá el arreglo?: ")
	fmt.Scan(&n)

	arreglo := make([]int, n)

	fmt.Println("Ingrese los elementos del arreglo:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&arr[i])
	}

	mayor := arreglo[0]

	for i := 1; i < n; i++ {
		if arreglo[i] > mayor {
			mayor = arreglo[i]
		}
	}

	fmt.Println("El mayor elemento del arreglo es:", mayor)
}
*/
