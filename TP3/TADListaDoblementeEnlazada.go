package main

// Estructura de nodo para lista simple
type Nodo struct {
	valor     int
	siguiente *Nodo
}

type Lista struct {
	cabeza *Nodo
	tamaño int
}

// Observadores básicos
func (l *Lista) verTamaño() int {
	return l.tamaño // O(1)
}

func (l *Lista) esVacia() bool {
	return l.tamaño == 0 // O(1)
}

func (l *Lista) primero() int {
	if l.esVacia() {
		panic("La lista está vacía")
	}
	return l.cabeza.valor // O(1)
}

func (l *Lista) pertenece(valor int) bool {
	aux := l.cabeza
	for aux != nil {
		if aux.valor == valor {
			return true
		}
		aux = aux.siguiente
	}
	return false // O(n)
}

// Generadores
func listaVacia() *Lista {
	nuevaLista := new(Lista)
	nuevaLista.cabeza = nil
	nuevaLista.tamaño = 0
	return nuevaLista // O(1)
}

func a_partir_de_secuencia_lista(secuencia []int) *Lista {
	lista := listaVacia()
	for _, valor := range secuencia {
		lista.agregarFinal(valor)
	}
	return lista // O(n)
}

// Otras operaciones
func (l *Lista) agregarInicio(valor int) {
	nuevo := &Nodo{valor: valor, siguiente: l.cabeza}
	l.cabeza = nuevo
	l.tamaño++
} // O(1)

func (l *Lista) agregarFinal(valor int) {
	nuevo := &Nodo{valor: valor, siguiente: nil}
	if l.esVacia() {
		l.cabeza = nuevo
	} else {
		aux := l.cabeza
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		aux.siguiente = nuevo
	}
	l.tamaño++
} // O(n)

func (l *Lista) eliminarInicio() int {
	if l.esVacia() {
		panic("La lista está vacía")
	}
	valor := l.cabeza.valor
	l.cabeza = l.cabeza.siguiente
	l.tamaño--
	return valor
} // O(1)

func (l *Lista) vaciar() {
	l.cabeza = nil
	l.tamaño = 0
} // O(1)
