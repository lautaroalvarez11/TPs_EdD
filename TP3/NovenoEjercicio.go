package main

import "fmt"

// Nodo representa un nodo doblemente enlazado
type Nodo struct {
	valor int
	prev  *Nodo
	next  *Nodo
}

// Deque representa la cola doblemente terminada
type Deque struct {
	head *Nodo // frente
	tail *Nodo // final
}

// Encolar al final (cola)
func (d *Deque) EncolarFinal(valor int) {
	nuevo := &Nodo{valor: valor}
	if d.tail == nil {
		d.head = nuevo
		d.tail = nuevo
	} else {
		d.tail.next = nuevo
		nuevo.prev = d.tail
		d.tail = nuevo
	}
}

// Encolar al frente (como una pila)
func (d *Deque) EncolarFrente(valor int) {
	nuevo := &Nodo{valor: valor}
	if d.head == nil {
		d.head = nuevo
		d.tail = nuevo
	} else {
		nuevo.next = d.head
		d.head.prev = nuevo
		d.head = nuevo
	}
}

// Desencolar del frente
func (d *Deque) DesencolarFrente() (int, bool) {
	if d.head == nil {
		return 0, false
	}
	valor := d.head.valor
	d.head = d.head.next
	if d.head != nil {
		d.head.prev = nil
	} else {
		d.tail = nil // cola vacía
	}
	return valor, true
}

// Desencolar del final
func (d *Deque) DesencolarFinal() (int, bool) {
	if d.tail == nil {
		return 0, false
	}
	valor := d.tail.valor
	d.tail = d.tail.prev
	if d.tail != nil {
		d.tail.next = nil
	} else {
		d.head = nil // cola vacía
	}
	return valor, true
}

// Mostrar la cola de izquierda a derecha
func (d *Deque) Mostrar() {
	for actual := d.head; actual != nil; actual = actual.next {
		fmt.Printf("%d ", actual.valor)
	}
	fmt.Println()
}

func main() {
	deque := Deque{}

	// Encolamos por ambos lados
	deque.EncolarFinal(10)
	deque.EncolarFinal(20)
	deque.EncolarFrente(5)
	deque.EncolarFrente(1)

	fmt.Print("Deque actual: ")
	deque.Mostrar()

	// Desencolamos
	valor, _ := deque.DesencolarFrente()
	fmt.Println("Desencolado del frente:", valor)

	valor, _ = deque.DesencolarFinal()
	fmt.Println("Desencolado del final:", valor)

	fmt.Print("Deque actual: ")
	deque.Mostrar()
}
