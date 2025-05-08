package main

import "fmt"

type Node struct {
	valor     int
	izquierda *Node
	derecha   *Node
}

type ABB struct {
	raiz *Node
}

// Ejercicio e - Función para calcular la cantidad total de nodos de un ABB

func (abb *ABB) ContarNodos() int {
	return contarNodos(abb.raiz)
}

// Función recursiva auxiliar
func contarNodos(nodo *Node) int {
	if nodo == nil {
		return 0
	}
	return 1 + contarNodos(nodo.izquierda) + contarNodos(nodo.derecha)
}

// Ejemplo de uso
func main() {
	abb := &ABB{}
	abb.Insertar(20)
	abb.Insertar(10)
	abb.Insertar(5)
	abb.Insertar(30)
	abb.Insertar(25)

	fmt.Println("Cantidad total de nodos:", abb.ContarNodos())
}

// Salida -> Cantidad total de nodos: 5

/*
Paso a paso:
- contarNodos(raiz) llama recursivamente a los hijos izquierdo y derecho de cada nodo.

- Cada nodo suma 1 a la cuenta.

- En total, se recorren todos los nodos exactamente una vez → complejidad O(n).
*/
