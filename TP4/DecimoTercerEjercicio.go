package main

import (
	"fmt"
)

type Nodo struct {
	valor     int
	izquierda *Nodo
	derecha   *Nodo
	altura    int
}

type AVL struct {
	raiz *Nodo
}

func altura(n *Nodo) int {
	if n == nil {
		return 0
	}
	return n.altura
}

func actualizarAltura(n *Nodo) {
	n.altura = 1 + max(altura(n.izquierda), altura(n.derecha))
}

func factorEquilibrio(n *Nodo) int {
	if n == nil {
		return 0
	}
	return altura(n.derecha) - altura(n.izquierda)
}

func rotarIzquierda(z *Nodo) *Nodo {
	y := z.derecha
	T2 := y.izquierda
	y.izquierda = z
	z.derecha = T2

	actualizarAltura(z)
	actualizarAltura(y)
	return y
}

func rotarDerecha(z *Nodo) *Nodo {
	y := z.izquierda
	T3 := y.derecha
	y.derecha = z
	z.izquierda = T3

	actualizarAltura(z)
	actualizarAltura(y)
	return y
}

// Para punto B
func insertar(n *Nodo, valor int) *Nodo {
	if n == nil {
		return &Nodo{valor: valor, altura: 1}
	}

	if valor < n.valor {
		n.izquierda = insertar(n.izquierda, valor)
	} else if valor > n.valor {
		n.derecha = insertar(n.derecha, valor)
	} else {
		return n // no se permiten duplicados
	}

	actualizarAltura(n)
	fe := factorEquilibrio(n)

	// Rotar izquierda izquierda
	if fe < -1 && valor < n.izquierda.valor {
		return rotarDerecha(n)
	}
	// Rotar derecha derecha
	if fe > 1 && valor > n.derecha.valor {
		return rotarIzquierda(n)
	}
	// Rotar izquierda derecha
	if fe < -1 && valor > n.izquierda.valor {
		n.izquierda = rotarIzquierda(n.izquierda)
		return rotarDerecha(n)
	}
	// Rotar derecha izquierda
	if fe > 1 && valor < n.derecha.valor {
		n.derecha = rotarDerecha(n.derecha)
		return rotarIzquierda(n)
	}

	return n
}

// Para punto C
func buscarAVL(nodo *Nodo, valor int) bool {
	if nodo == nil {
		return false
	}
	if valor == nodo.valor {
		return true
	} else if valor < nodo.valor {
		return buscarAVL(nodo.izquierda, valor)
	} else {
		return buscarAVL(nodo.derecha, valor)
	}
}

func minimo(n *Nodo) *Nodo {
	actual := n
	for actual.izquierda != nil {
		actual = actual.izquierda
	}
	return actual
}

func mayor(m *Nodo) *Nodo {
	actual := m
	for actual.derecha != nil {
		actual = actual.derecha
	}
	return actual
}

func main() {
	var raiz *Nodo

	codigos := []int{
		1012,
		1020,
		1005,
		1030,
		1008,
		1015,
		1025,
	}

	// Insertar en AVL
	for _, codigos := range codigos {
		raiz = insertar(raiz, codigos)
	}

	fmt.Println("Códigos insertados en el AVL.")

	// Ejemplo de búsqueda
	buscar := 1015
	encontrado := buscarAVL(raiz, buscar)
	fmt.Printf("¿El código %d está en el inventario?: %v\n", buscar, encontrado)

	// Ejemplo de mínimo y máximo
	fmt.Println("Código menor:", minimo(raiz).valor)
	fmt.Println("Código mayor:", mayor(raiz).valor)
}

/*
Para punto D
En un ABB no balanceado (insertando en orden):
1012
   \
   1020
      \
      1030
     /
   1025
   /
 1015
 /
1008
/
1005

Comparaciones para encontrar 1025:
- 1025 > 1012 → derecha (1)
- 1025 > 1020 → derecha (2)
- 1025 < 1030 → izquierda (3)
- 1025 == 1025 → Encontrado (4)
- Total: 4 comparaciones

En un AVL:
        1012
       /     \
   1005       1020
     \       /    \
     1008  1015   1030
                   /
                1025

Comparaciones para encontrar 1025:
- 1025 > 1012 → derecha (1)
- 1025 > 1020 → derecha (2)
- 1025 < 1030 → izquierda (3)
- 1025 == 1025 → Encontrado (4)
- Total: 4 comparaciones

En este caso puntual, el número de comparaciones coincide porque el árbol no se volvió totalmente desbalanceado.
Pero si se insertara valores en orden creciente o decreciente, el ABB se degrada y podría tardar hasta n comparaciones.
Mientras que el AVL se mantiene en O(log n).
*/

/*
Para punto E
¿Por qué AVL es más adecuado?:
Para un sistema de inventario digital donde las búsquedas deben ser rápidas, constantes y confiables,
el AVL es más adecuado que el ABB no balanceado, porque garantiza un árbol equilibrado, evitando los peores casos y
manteniendo la eficiencia incluso si los códigos llegan en orden creciente, aleatorio o descendente.
*/
