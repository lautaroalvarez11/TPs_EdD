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

/*
Para punto D
func (a *ABB) buscarConComparaciones(valor int) (bool, int) {
	return buscarABBConContador(a.raiz, valor, 0)
}

func buscarABBConContador(nodo *Nodo, valor int, comp int) (bool, int) {
	if nodo == nil {
		return false, comp
	}
	comp++
	if valor == nodo.valor {
		return true, comp
	} else if valor < nodo.valor {
		return buscarABBConContador(nodo.izquierda, valor, comp)
	} else {
		return buscarABBConContador(nodo.derecha, valor, comp)
	}
}

func buscarAVLConContador(nodo *Nodo, valor int, comp int) (bool, int) {
	if nodo == nil {
		return false, comp
	}
	comp++
	if valor == nodo.valor {
		return true, comp
	} else if valor < nodo.valor {
		return buscarAVLConContador(nodo.izquierda, valor, comp)
	} else {
		return buscarAVLConContador(nodo.derecha, valor, comp)
	}
}

- Cada vez que comparamos valor == nodo.valor, valor < nodo.valor, o valor > nodo.valor, se cuenta como una comparación.
- Cuantas más comparaciones, menos eficiente es la búsqueda.

Resultados esperados de ambas funciones:
- En un ABB no balanceado (si los datos se insertan en orden creciente), puede tardar hasta n comparaciones (peor caso).
- En un AVL, al estar balanceado, las comparaciones son aproximadamente log₂(n).
*/

/*
Para punto E
| Característica               | ABB (no balanceado)                          | AVL (balanceado)                            |
| ---------------------------- | -------------------------------------------- | ------------------------------------------- |
| Búsqueda (mejor caso)        | O(log n)                                     | O(log n)                                    |
| Búsqueda (peor caso)         | O(n) ← si está desbalanceado                 | O(log n) garantizado                        |
| Inserción                    | Simple y rápida                              | Más compleja (puede requerir rotaciones)    |
| Mantenimiento del equilibrio | No                                           | Sí, mantiene equilibrio después de insertar |
| Ideal para...                | Estructuras pequeñas o de inserción ordenada | Grandes volúmenes, consultas frecuentes     |

¿Por qué elegir AVL en este caso?:
1. Búsqueda rápida garantizada:
- En un sistema de biblioteca, las búsquedas (por ISBN) son muy frecuentes y críticas.
- Un ABB desbalanceado puede volverse una lista lineal, degradando el rendimiento a O(n).

2. Orden aleatorio de inserciones:
- No se puede controlar el orden en que llegan los ISBN.
- El AVL se ajusta automáticamente para mantener eficiencia.

3. Escalabilidad:
- A medida que la biblioteca crece, un ABB pierde eficiencia, mientras que un AVL mantiene rendimiento constante.

4. Consistencia en el tiempo de respuesta:
- Ideal para aplicaciones en producción donde se requiere velocidad constante y predecible.
*/

func main() {
	var raiz *Nodo

	isbns := []int{
		9780345391803,
		9780131103627,
		9780321573513,
		9780596009205,
		9780262033848,
	}

	// Insertar en AVL
	for _, isbn := range isbns {
		raiz = insertar(raiz, isbn)
	}

	fmt.Println("ISBNs insertados en el AVL.")

	// Ejemplo de búsqueda
	buscar := 9780596009205
	fmt.Printf("¿Está el ISBN %d en el AVL? %v\n", buscar, buscarAVL(raiz, buscar))
}
