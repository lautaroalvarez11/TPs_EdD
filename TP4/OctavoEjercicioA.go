package main

type Node struct {
	valor     int
	izquierda *Node
	derecha   *Node
}

type ABB struct {
	raiz *Node
}

// Ejercicio a - Función para insertar un valor en un ABB

func (abb *ABB) Insertar(valor int) {
	abb.raiz = insertarNodo(abb.raiz, valor)
}

func insertarNodo(nodo *Node, valor int) *Node {
	// Si el nodo es nil, crear un nuevo nodo con el valor
	if nodo == nil {
		return &Node{valor: valor}
	}

	// Si el valor es menor que el nodo actual, ir a la izquierda
	if valor < nodo.valor {
		nodo.izquierda = insertarNodo(nodo.izquierda, valor)
	} else if valor > nodo.valor { // Si el valor es mayor, ir a la derecha
		nodo.derecha = insertarNodo(nodo.derecha, valor)
	}
	return nodo // Retornar el nodo sin cambios
}

// Ejemplo de uso
func main() {

	abb := &ABB{}
	abb.Insertar(10)
	abb.Insertar(5)
	abb.Insertar(15)
	abb.Insertar(3)

}

/*
Paso a paso:
1. Creación del árbol binario de búsqueda (ABB):
	- Se crea una instancia del árbol binario de búsqueda (ABB) inicializando su raíz como nil (vacío).

2. Inserción del valor 10:
	- Se llama al método Insertar del ABB con el valor 10.
	- Como la raíz del ABB es nil, se crea un nuevo nodo con el valor 10 y se asigna como la raíz del árbol.
	- Estado del ABB:
		10

3. Inserción del valor 5:
	- Se llama al método Insertar del ABB con el valor 5.
	- La función insertarNodo compara 5 con el valor de la raíz (10):
		- Como 5 < 10, se mueve al subárbol izquierdo.
		- El subárbol izquierdo está vacío (nil), por lo que se crea un nuevo nodo con el valor 5 y se asigna como hijo izquierdo de la raíz.
	- Estado del ABB:
		   10
		  /
		 5

4. Inserción del valor 15:
	- Se llama al método Insertar con el valor 15.
	- La función insertarNodo compara 15 con el valor de la raíz (10):
		- Como 15 > 10, se mueve al subárbol derecho.
		- El subárbol derecho está vacío (nil), por lo que se crea un nuevo nodo con el valor 15 y se asigna como hijo derecho de la raíz.
	- Estado del ABB:
		   10
		  /  \
		 5   15

5. Inserción del valor 3:
	- Se llama al método Insertar con el valor 3.
	- La función insertarNodo compara 3 con el valor de la raíz (10):
		- Como 3 < 10, se mueve al subárbol izquierdo.
	- Luego, compara 3 con el valor del nodo izquierdo (5):
		- Como 3 < 5, se mueve al subárbol izquierdo de este nodo.
		- El subárbol izquierdo de 5 está vacío (nil), por lo que se crea un nuevo nodo con el valor 3 y se asigna como hijo izquierdo de 5.
	- Estado final del ABB:
		   10
		  /  \
		 5   15
		/
	   3
*/
