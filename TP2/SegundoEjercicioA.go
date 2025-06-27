package main

import "fmt"

/*
	Cuenta regresiva: Recibe un número entero e imprime todos los números comprendidos
	entre el mismo y el 0
*/

// FORMA RECURSIVA: O(n)

func imprimirNumeros(n int) {
	// Caso base: cuando n es 0, se imprime y se termina la recursión
	fmt.Println(n)
	if n == 0 {
		return
	}

	// Llamada recursiva: si n es positivo, restar 1; si es negativo, sumar 1
	if n > 0 {
		imprimirNumeros(n - 1) // Llamado recursivo con n-1
	} else {
		imprimirNumeros(n + 1) // Llamado recursivo con n+1
	}
}

func main() {
	var num int // Declarar una variable entera para almacenar la entrada del usuario

	// Pedir al usuario que ingrese un número entero
	fmt.Print("Ingrese un número entero: ")
	fmt.Scan(&num) // Leer el número ingresado y almacenarlo en la variable num

	imprimirNumeros(num) // Llamar a la función imprimirNumeros con el número ingresado por el usuario
}

/*
FORMA ITERATIVA: O(n)

func main() {
	var num int
	fmt.Print("Ingrese un número entero: ")
	fmt.Scan(&num)

	for i := num; i >= 0; i-- {
		fmt.Println(i)
	}
}
*/
