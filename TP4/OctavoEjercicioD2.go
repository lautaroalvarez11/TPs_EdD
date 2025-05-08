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

// Ejercicio d.2 - Función para encontrar el mínimo de un ABB en forma iterativa

func (abb *ABB) MinimoIterativo() (int, bool) {
	actual := abb.raiz

	if actual == nil {
		return 0, false // Árbol vacío
	}

	for actual.izquierda != nil {
		actual = actual.izquierda
	}

	return actual.valor, true
}

// Ejemplo de uso
func main() {
	abb := &ABB{}
	abb.Insertar(20)
	abb.Insertar(10)
	abb.Insertar(5)
	abb.Insertar(30)

	minIter, ok := abb.MinimoIterativo()
	if ok {
		fmt.Println("Mínimo (iterativo):", minIter)
	} else {
		fmt.Println("Árbol vacío (iterativo)")
	}
}

/*
| Método    | Ventaja                                 | Costo                                   |
| --------- | --------------------------------------- | --------------------------------------- |
| Recursivo | Más simple y directo                    | Usa pila de llamadas (consumo de stack) |
| Iterativo | Más eficiente en memoria (no usa stack) | Ligeramente más largo de escribir       |
*/
