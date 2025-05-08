package main

// Estructura de nodo para lista doble
type NodoDoble struct {
	valor     int
	previo    *NodoDoble
	siguiente *NodoDoble
}

type ListaDoble struct {
	cabeza *NodoDoble
	cola   *NodoDoble
	tamaño int
}

// Observadores básicos
func (ld *ListaDoble) verTamaño() int {
	return ld.tamaño // O(1)
}

func (ld *ListaDoble) esVacia() bool {
	return ld.tamaño == 0 // O(1)
}

func (ld *ListaDoble) primero() int {
	if ld.esVacia() {
		panic("La lista está vacía")
	}
	return ld.cabeza.valor // O(1)
}

func (ld *ListaDoble) ultimo() int {
	if ld.esVacia() {
		panic("La lista está vacía")
	}
	return ld.cola.valor // O(1)
}

func (ld *ListaDoble) pertenece(valor int) bool {
	aux := ld.cabeza
	for aux != nil {
		if aux.valor == valor {
			return true
		}
		aux = aux.siguiente
	}
	return false // O(n)
}

// Generadores
func listaDobleVacia() *ListaDoble {
	nuevaLista := new(ListaDoble)
	nuevaLista.cabeza = nil
	nuevaLista.cola = nil
	nuevaLista.tamaño = 0
	return nuevaLista // O(1)
}

func a_partir_de_secuencia_listaDoble(secuencia []int) *ListaDoble {
	lista := listaDobleVacia()
	for _, valor := range secuencia {
		lista.agregarFinal(valor)
	}
	return lista // O(n)
}

// Otras operaciones
func (ld *ListaDoble) agregarInicio(valor int) {
	nuevo := &NodoDoble{valor: valor, siguiente: ld.cabeza, previo: nil}
	if ld.esVacia() {
		ld.cabeza = nuevo
		ld.cola = nuevo
	} else {
		ld.cabeza.previo = nuevo
		ld.cabeza = nuevo
	}
	ld.tamaño++
} // O(1)

func (ld *ListaDoble) agregarFinal(valor int) {
	nuevo := &NodoDoble{valor: valor, siguiente: nil, previo: ld.cola}
	if ld.esVacia() {
		ld.cabeza = nuevo
		ld.cola = nuevo
	} else {
		ld.cola.siguiente = nuevo
		ld.cola = nuevo
	}
	ld.tamaño++
} // O(1)

func (ld *ListaDoble) eliminarInicio() int {
	if ld.esVacia() {
		panic("La lista está vacía")
	}
	valor := ld.cabeza.valor
	ld.cabeza = ld.cabeza.siguiente
	if ld.cabeza != nil {
		ld.cabeza.previo = nil
	} else {
		ld.cola = nil // La lista quedó vacía
	}
	ld.tamaño--
	return valor
} // O(1)

func (ld *ListaDoble) eliminarFinal() int {
	if ld.esVacia() {
		panic("La lista está vacía")
	}
	valor := ld.cola.valor
	ld.cola = ld.cola.previo
	if ld.cola != nil {
		ld.cola.siguiente = nil
	} else {
		ld.cabeza = nil // La lista quedó vacía
	}
	ld.tamaño--
	return valor
} // O(1)

func (ld *ListaDoble) vaciar() {
	ld.cabeza = nil
	ld.cola = nil
	ld.tamaño = 0
} // O(1)
