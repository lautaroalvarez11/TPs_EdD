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

// Ejercicio c - Implementar la función eliminar(nodo, valor) -> nodo para eliminar un valor del ABB

func (abb *ABB) Eliminar(valor int) {
	abb.raiz = eliminarNodo(abb.raiz, valor)
}

func eliminarNodo(nodo *Node, valor int) *Node {
	// Caso base: si el nodo es nil, valor no encontrado
	if nodo == nil {
		fmt.Println("El elemento a eliminar no existe en el árbol")
		return nil // Valor no encontrado
	}

	if valor < nodo.valor {
		nodo.izquierda = eliminarNodo(nodo.izquierda, valor)
	} else if valor > nodo.valor {
		nodo.derecha = eliminarNodo(nodo.derecha, valor)
	} else {
		// Caso 1: Nodo sin hijos
		if nodo.izquierda == nil && nodo.derecha == nil {
			return nil
		}

		// Caso 2: Nodo con un solo hijo
		if nodo.izquierda == nil {
			return nodo.derecha
		}
		if nodo.derecha == nil {
			return nodo.izquierda
		}

		// Caso 3: Nodo con dos hijos, buscar el sucesor inorden (mínimo del subárbol derecho)
		minimo := encontrarMinimo(nodo.derecha)
		nodo.valor = minimo.valor
		// Eliminar el sucesor del subárbol derecho
		nodo.derecha = eliminarNodo(nodo.derecha, minimo.valor)
	}

	return nodo
}

// Función auxiliar: encontrar el mínimo nodo (usada en el caso 3)
func encontrarMinimo(nodo *Node) *Node {
	if nodo == nil {
		return nil
	}
	for nodo.izquierda != nil {
		nodo = nodo.izquierda
	}
	return nodo
}

// Ejemplo de uso

func main() {
	abb := &ABB{}
	abb.Insertar(20)
	abb.Insertar(10)
	abb.Insertar(30)
	abb.Insertar(25)
	abb.Insertar(35)

	fmt.Println("Eliminando 30...")
	abb.Eliminar(30) // Nodo con dos hijos

	fmt.Println("Eliminando 10...")
	abb.Eliminar(10) // Nodo con un hijo

	fmt.Println("Eliminando 35...")
	abb.Eliminar(35) // Nodo hoja
}
