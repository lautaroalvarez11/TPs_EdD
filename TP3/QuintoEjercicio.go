/*
- TAD Cola
- Igualdad Observacional: Dos colas son observacionalmente iguales si tienen los mismos elementos en el mismo orden y la misma longitud.
- Usa: Bool, Nat, a (elemento genérico)
- Parámetro formal: a
- Género: Cola<a>
- Observadores básicos:
	•	es_vacia(cola: Cola<a>) -> Bool
	•	frente(cola: Cola<a>) -> a
	•	longitud(cola: Cola<a>) -> Nat
- Generadores:
	•	vacia() -> Cola<a>
	•	encolar(cola: Cola<a>, elemento: a) -> Cola<a>
- Otras operaciones:
	•	desencolar(cola: Cola<a>) -> Cola<a>
	•	pertenece(elemento: a, cola: Cola<a>) -> Bool
- Axiomas:
	•	es_vacia(vacia()) -> true
	•	es_vacia(encolar(cola, elemento)) -> false
	•	frente(encolar(vacia(), elemento)) -> e
	•	frente(encolar(encolar(cola, x), elemento)) -> frente(encolar(cola, x))
	•	longitud(vacia()) -> 0
	•	longitud(encolar(cola, elemento)) -> 1 + longitud(cola)
	•	desencolar(encolar(vacia(), elemento)) -> vacia()
	•	desencolar(encolar(cola, elemento)) -> encolar(desencolar(cola), elemento) si c ≠ vacia()
	•	pertenece(elemento, vacia()) -> false
	•	pertenece(elemento, encolar(cola, x)) -> (elemento = frente(encolar(cola, x))) or pertenece(elemento, cola)
- Exporta: Cola<a>, vacia, encolar, desencolar, frente, es_vacia, longitud, pertenece

Implementación del TAD Cola: Utilicé una estructura con nodos (listas enlazadas) porque tiene acceso eficiente al frente y al final,
además de que permite agregar y quitar elementos en O(1).

Justificación de la estructura interna: Al utilizar nodos enlazados se evitan copias grandes, el acceso al frente y la cola es inmediato,
no se necesita conocer el tamaño por adelantado.
*/

package cola

type Nodo[a any] struct {
	dato      a
	siguiente *Nodo[a]
}

type Cola[a any] struct {
	frente   *Nodo[a]
	fin      *Nodo[a]
	longitud int
}

func vacia[a any]() Cola[a] {
	return Cola[a]{nil, nil, 0}
}

func esVacia[a any](cola Cola[a]) bool {
	return cola.frente == nil
}

func encolar[a any](cola *Cola[a], elemento a) {
	nuevo := &Nodo[a]{dato: elemento}
	if cola.fin != nil {
		cola.fin.siguiente = nuevo
	} else {
		cola.frente = nuevo
	}
	cola.fin = nuevo
	cola.longitud++
}

func Desencolar[a any](cola *Cola[a]) {
	if cola.frente != nil {
		cola.frente = cola.frente.siguiente
		cola.longitud--
		if cola.frente == nil {
			cola.fin = nil
		}
	}
}

func Frente[a any](cola Cola[a]) a {
	return cola.frente.dato
}

func Longitud[a any](cola Cola[a]) int {
	return cola.longitud
}

func Pertenece[a comparable](cola Cola[a], elemento a) bool {
	aux := cola.frente
	for aux != nil {
		if aux.dato == elemento {
			return true
		}
		aux = aux.siguiente
	}
	return false
}

/*
Complejidad de operaciones:

Operación	Complejidad   Justificación

EsVacia		    O(1)      Revisa un puntero.
Encolar		    O(1)      Accede directo al fin.
Desencolar      O(1)      Modifica el puntero frente.
Frente          O(1)      Retorna un campo.
Longitud        O(1)      Campo guardado.
Pertenece       O(n)      Recorre la cola entera en el peor caso.
*/
