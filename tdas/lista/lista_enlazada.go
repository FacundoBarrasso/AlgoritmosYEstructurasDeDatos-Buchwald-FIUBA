package lista

/* Definición de las estructuras de la lista enlazada */

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] {
	return new(listaEnlazada[T])
}

/* Definición de las primitivas de la lista enlazada */

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevoNodo := crearNodoLista(dato)

	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	} else {
		nuevoNodo.siguiente = lista.primero
	}

	lista.primero = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevoNodo := crearNodoLista(dato)

	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}

	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	panicListaVacia(lista)

	nodoBorrado := lista.primero.dato
	lista.primero = lista.primero.siguiente
	lista.largo--

	if lista.largo == 0 {
		lista.ultimo = nil
	}

	return nodoBorrado
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	panicListaVacia(lista)
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	panicListaVacia(lista)
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

// crearNodoLista crea un nuevo nodo de lista con el dato pasado por parámetro.
func crearNodoLista[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{dato, nil}
}

// panicListaVacia lanza un pánico si la lista está vacía.
func panicListaVacia[T any](lista *listaEnlazada[T]) {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	for nodo := lista.primero; nodo != nil; nodo = nodo.siguiente {
		if !visitar(nodo.dato) {
			break
		}
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{lista, lista.primero, nil}
}

/* Definición de las estructuras del iterador externo */

type iterListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

/* Definición de las primitivas del iterador externo */

func (iter *iterListaEnlazada[T]) VerActual() T {
	panicIter(iter)
	return iter.actual.dato
}

func (iter *iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterListaEnlazada[T]) Siguiente() {
	panicIter(iter)
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iterListaEnlazada[T]) Insertar(dato T) {
	elem := crearNodoLista(dato)

	if iter.anterior == nil {
		elem.siguiente = iter.lista.primero
		iter.lista.primero = elem
	} else {
		elem.siguiente = iter.anterior.siguiente
		iter.anterior.siguiente = elem
	}

	if iter.actual == nil {
		iter.lista.ultimo = elem
	}

	iter.actual = elem
	iter.lista.largo++
}

func (iter *iterListaEnlazada[T]) Borrar() T {
	panicIter(iter)

	borrado := iter.actual.dato

	if iter.anterior == nil {
		iter.lista.primero = iter.actual.siguiente
	}
	if iter.actual.siguiente == nil {
		iter.lista.ultimo = iter.anterior
	}

	iter.actual = iter.actual.siguiente

	if iter.anterior != nil {
		iter.anterior.siguiente = iter.actual
	}

	iter.lista.largo--
	return borrado
}

// panicIter lanza un pánico si el iterador terminó de iterar.
func panicIter[T any](iter *iterListaEnlazada[T]) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}
