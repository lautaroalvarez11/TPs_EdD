package main

import (
	"fmt"
)

type Node struct {
	valor     int
	izquierda *Node
	derecha   *Node
}

type ABB struct {
	raiz *Node
}

// Ejercicio d.1 - Función para encontrar el mínimo de un ABB en forma recursiva

func (abb *ABB) MinimoRecursivo() (int, bool) {
	// Verificar si el árbol está vacío
	nodo := minimoRecursivo(abb.raiz)

	if nodo == nil { // Si el nodo es nil, significa que el árbol está vacío
		return 0, false // Árbol vacío
	}
	return nodo.valor, true // Retornar el valor del nodo mínimo y true indicando que se encontró
}

func minimoRecursivo(nodo *Node) *Node {
	// Caso base: si el nodo es nil, retornar nil
	if nodo == nil || nodo.izquierda == nil {
		return nodo
	}
	// Caso recursivo: si el nodo no es nil, seguir buscando en la izquierda
	return minimoRecursivo(nodo.izquierda)
}

// Ejemplo de uso
func main() {
	abb := &ABB{}
	abb.Insertar(20)
	abb.Insertar(10)
	abb.Insertar(5)
	abb.Insertar(30)

	minRec, ok := abb.MinimoRecursivo()
	if ok {
		fmt.Println("Mínimo (recursivo):", minRec)
	} else {
		fmt.Println("Árbol vacío (recursivo)")
	}
}
