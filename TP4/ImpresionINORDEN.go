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

// Imprimir el ABB en inorden (izquierda - raíz - derecha)
func (abb *ABB) ImprimirInorden() {
	fmt.Print("Árbol inorden: ")
	imprimirInorden(abb.raiz)
	fmt.Println()
}

func imprimirInorden(nodo *Node) {
	if nodo != nil {
		imprimirInorden(nodo.izquierda)
		fmt.Printf("%d ", nodo.valor)
		imprimirInorden(nodo.derecha)
	}
}

// Ejemplo de uso
func main() {
	abb := &ABB{}
	abb.Insertar(20)
	abb.Insertar(10)
	abb.Insertar(30)
	abb.Insertar(25)
	abb.Insertar(35)

	fmt.Println("Antes de eliminar:")
	abb.ImprimirInorden()

	abb.Eliminar(30)
	fmt.Println("Después de eliminar 30:")
	abb.ImprimirInorden()

	abb.Eliminar(10)
	fmt.Println("Después de eliminar 10:")
	abb.ImprimirInorden()
}

/*
Paso a paso:
- Antes de eliminar:
	- Árbol inorden: 10 20 25 30 35
- Después de eliminar 30:
	- Árbol inorden: 10 20 25 35
- Después de eliminar 10:
	- Árbol inorden: 20 25 35
*/
