package lista

type Lista[T any] interface {

	// EstaVacia devuelve true si la lista no tiene ningun elemento, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento al principio de la lista. Si esta vacía el nuevo elemento será el
	// primero y el último.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento al final de la lista. Si esta vacía el nuevo elemento será el
	// primero y el último.
	InsertarUltimo(T)

	// BorrarPrimero elimina el primer elemento de la lista y lo devuelve. Si la lista esta vacía, entra
	// en pánico con el mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero devuelve el primer elemento de la lista. Si la lista esta vacía, entra en pánico con el
	// mensaje "La lista esta vacia".
	VerPrimero() T

	// VerUltimo devuelve el último elemento de la lista. Si la lista esta vacía, entra en pánico con el
	// mensaje "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos que tiene la lista.
	Largo() int

	// Iterar recorre la lista y llama a la función visitar con cada elemento. Si visitar devuelve false, la
	// iteración se detiene.
	Iterar(visitar func(T) bool)

	// Iterador crea un nuevo iterador para la lista.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual devuelve el elemento al que está apuntando el iterador. Si el iterador ya iteró todos los elementos,
	//entra en pánico con un mensaje "El iterador termino de iterar".
	VerActual() T

	// HaySiguiente devuelve true si hay por lo menos un elemento sin iterar, false en caso contrario.
	HaySiguiente() bool

	// Siguiente avanza al siguiente elemento de la lista. Si el iterador ya iteró todos los elementos, entra en pánico
	//con un mensaje "El iterador termino de iterar".
	Siguiente()

	// Insertar inserta un nuevo dato entre el elemento actual y el anterior. El iterador apunta al elemento insertado.
	// Si la lista tenía únicamente un elemento, el dato insertado pasará a ser el primer elemento.
	Insertar(T)

	// Borrar elimina el elemento de la lista al que está apuntando el iterador y este apunta al elemento siguiente.
	// Devuelve el elemento borrado. Si el iterador ya iteró todos los elementos, entra en pánico con un mensaje "El iterador termino de iterar".
	Borrar() T
}
