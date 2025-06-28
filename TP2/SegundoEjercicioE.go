package main

import "fmt"

/*
	Invertir: Dado un arreglo de enteros, invertirlo.
*/

// FORMA RECURSIVA: O(n)

func invertir(arr []int, inicio int, fin int) {
	// Caso base: Cuando los índices se cruzan o se tocan, ya está invertido
	if inicio >= fin {
		return
	}

	// Intercambiar los elementos en las posiciones 'inicio' y 'fin'
	arr[inicio], arr[fin] = arr[fin], arr[inicio]

	// Llamada recursiva moviendo los índices hacia el centro
	invertir(arr, inicio+1, fin-1)
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

	invertir(arr, 0, len(arr)-1)

	fmt.Println("Arreglo invertido:")
	fmt.Println(arr)
}

/*
FORMA ITERATIVA: O(n)

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

	arregloInvertido := make([]int, n)

	for j := 0; j < n; j++ {
		arregloInvertido[j] = arr[n-j-1]
	}

	fmt.Println("Arreglo invertido:")
	fmt.Println(arregloInvertido)
}
*/
