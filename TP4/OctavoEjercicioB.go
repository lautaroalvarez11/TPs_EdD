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

// Ejercicio b - Función que busque un valor en un ABB y devuelva un booleano indicando si el valor está presente o no

func (abb *ABB) Buscar(valor int) (bool, int) {
	return buscarNodo(abb.raiz, valor, 0)
}

func buscarNodo(nodo *Node, valor int, comparaciones int) (bool, int) {
	// Caso base: si el nodo es nil, el valor no está presente y se devuelve false junto con el número de comparaciones realizadas
	if nodo == nil {
		fmt.Println("El elemento no existe en el árbol")
		return false, comparaciones
	}

	// Cada vez que se compara el valor buscado con el nodo actual, se incrementa el contador de comparaciones
	comparaciones++ // Contar una comparación con el nodo actual

	// Comparar el valor buscado con el valor del nodo actual
	if valor == nodo.valor { // Si el valor es igual
		return true, comparaciones // Se devuelve true y el número de comparaciones realizadas
	} else if valor < nodo.valor { // Si el valor buscado es menor que el valor del nodo actual
		return buscarNodo(nodo.izquierda, valor, comparaciones) // Se busca en el subárbol izquierdo
	} else { // Si el valor buscado es mayor que el valor del nodo actual
		return buscarNodo(nodo.derecha, valor, comparaciones) // Se busca en el subárbol derecho
	}
}

// Ejemplo de uso

func main() {
	abb := &ABB{}
	abb.Insertar(10)
	abb.Insertar(5)
	abb.Insertar(15)
	abb.Insertar(3)

	// Buscar un valor
	encontrado, comparaciones := abb.Buscar(3)
	fmt.Printf("¿Está el valor 3?: %v (comparaciones: %d)\n", encontrado, comparaciones)

	encontrado, comparaciones = abb.Buscar(20)
	fmt.Printf("¿Está el valor 20?: %v (comparaciones: %d)\n", encontrado, comparaciones)
}

/*
Paso a paso:
- Compara con 10 → 3 < 10 → va a la izquierda (1ª comparación)

- Compara con 5 → 3 < 5 → va a la izquierda (2ª comparación)

- Compara con 3 → Encontrado (3ª comparación)
*/

/*
Análisis del peor caso de comparaciones:
- En un ABB desbalanceado (como una lista), el peor caso es recorrer todos los nodos:
	→ Comparaciones = número de niveles = altura del árbol = O(n)

- En un ABB balanceado:
	→ Comparaciones = O(log n)
*/
