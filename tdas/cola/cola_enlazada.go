package cola

type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

type nodo[T any] struct {
	dato T
	prox *nodo[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{nil, nil}
}

// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario.
func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

// VerPrimero obtiene el valor del primero de la cola. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".
func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.primero.dato
}

// Encolar agrega un nuevo elemento a la cola, al final de la misma.
func (c *colaEnlazada[T]) Encolar(dato T) {
	if c.EstaVacia() {
		c.primero = crearNodo(dato)
		c.ultimo = c.primero
		return
	} else {
		c.ultimo.prox = crearNodo(dato)
		c.ultimo = c.ultimo.prox
	}
}

// Desencolar saca el primer elemento de la cola. Si la cola tiene elementos, se quita el primero de la misma, y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".
func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := c.primero.dato
	c.primero = c.primero.prox
	if c.primero == nil {
		c.ultimo = nil
	}
	return dato
}

func crearNodo[T any](dato T) *nodo[T] {
	return &nodo[T]{dato, nil}
}
