package main

// Estructura de datos para la cola
type Nodo struct {
	valor     int
	siguiente *Nodo
}

type Cola struct {
	cabeza *Nodo
	cola   *Nodo
	tamaño int
}

// Observadores basicos
func (c *Cola) verTamaño() int {
	return c.tamaño // O(1)
}

func (c *Cola) esVacia() bool {
	return c.tamaño == 0 // O(1)
}

func (c *Cola) primero() int {
	if c.esVacia() {
		panic("La cola está vacía")
	}
	return c.cabeza.valor // O(1)
}

func (c *Cola) ultimo() int {
	if c.esVacia() {
		panic("La cola está vacía")
	}
	return c.cola.valor // O(1)
}

func (c *Cola) pertenece(valor int) bool {
	aux := c.cabeza
	for aux != nil {
		if aux.valor == valor {
			return true
		}
		aux = aux.siguiente
	}
	return false // O(n)
}

// Generadores
func vacio() *Cola {
	nuevaCola := new(Cola)
	nuevaCola.cabeza = nil
	nuevaCola.cola = nil
	nuevaCola.tamaño = 0
	return nuevaCola // O(1)
}
func a_partir_de_secuencia(secuencia []int) *Cola {
	cola := vacio()
	for _, valor := range secuencia {
		cola.encolar(valor)
	}
	return cola // O(n)
}

// Otras operaciones
func (c *Cola) encolar(valor int) {
	if c.esVacia() {
		c.cabeza = &Nodo{valor: valor, siguiente: nil}
		c.cola = c.cabeza
	} else {
		c.cola.siguiente = &Nodo{valor: valor, siguiente: nil}
		c.cola = c.cola.siguiente
	}
	c.tamaño++ // O(1)
}

func (c *Cola) desencolar() int {
	if c.esVacia() {
		panic("La cola está vacía")
	}
	valor := c.cabeza.valor
	c.cabeza = c.cabeza.siguiente
	c.tamaño--
	if c.esVacia() {
		c.cola = nil
	}
	return valor // O(1)
}

func (c *Cola) vaciar() {
	c.cabeza = nil
	c.cola = nil
	c.tamaño = 0
} // O(1)
