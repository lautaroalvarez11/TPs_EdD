package main

// Estructura de datos para la pila
type Nodo struct {
	valor     int
	siguiente *Nodo
}

type Pila struct {
	tope   *Nodo
	tamaño int
}

// Observadores básicos
func (p *Pila) verTamaño() int {
	return p.tamaño // O(1)
}

func (p *Pila) esVacia() bool {
	return p.tamaño == 0 // O(1)
}

func (p *Pila) verTope() int {
	if p.esVacia() {
		panic("La pila está vacía")
	}
	return p.tope.valor // O(1)
}

func (p *Pila) pertenece(valor int) bool {
	aux := p.tope
	for aux != nil {
		if aux.valor == valor {
			return true
		}
		aux = aux.siguiente
	}
	return false // O(n)
}

// Generadores
func vacia() *Pila {
	nuevaPila := new(Pila)
	nuevaPila.tope = nil
	nuevaPila.tamaño = 0
	return nuevaPila // O(1)
}

func a_partir_de_secuencia(secuencia []int) *Pila {
	pila := vacia()
	for i := len(secuencia) - 1; i >= 0; i-- {
		pila.apilar(secuencia[i])
	}
	return pila // O(n)
}

// Otras operaciones
func (p *Pila) apilar(valor int) {
	nuevo := &Nodo{valor: valor, siguiente: p.tope}
	p.tope = nuevo
	p.tamaño++
} // O(1)

func (p *Pila) desapilar() int {
	if p.esVacia() {
		panic("La pila está vacía")
	}
	valor := p.tope.valor
	p.tope = p.tope.siguiente
	p.tamaño--
	return valor
} // O(1)

func (p *Pila) vaciar() {
	p.tope = nil
	p.tamaño = 0
} // O(1)
