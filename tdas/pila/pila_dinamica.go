package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

const VALOR_PARA_AGRANDAR = 2
const VALOR_PARA_ACHICAR = 2

// CrearPilaDinamica crea una nueva pila dinámica vacía.
func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, 0), cantidad: 0}
}

// EstaVacia devuelve verdadero si la pila no tiene elementos apilados, false en caso contrario.
func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

// VerTope obtiene el valor del tope de la pila. Si la pila tiene elementos se devuelve el valor del tope.
// Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

// Apilar agrega un nuevo elemento a la pila sin utilizar append y redimensionando solo si se llega al tope.
func (p *pilaDinamica[T]) Apilar(dato T) {
	if p.cantidad == len(p.datos) {
		p.redimensionar("arriba")
	}
	p.datos[p.cantidad] = dato
	p.cantidad++
}

// Desapilar saca el elemento tope de la pila. Si la pila tiene elementos, se quita el tope de la pila, y
// se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	p.cantidad--
	if p.cantidad*4 <= len(p.datos) {
		p.redimensionar("abajo")
	}
	return p.datos[p.cantidad]
}

// redimensionar redimensiona el slice de datos de la pila según la dirección indicada. Si la dirección es "arriba", duplica la capacidad.
// Si la dirección es "abajo", reduce la capacidad a la mitad.
func (p *pilaDinamica[T]) redimensionar(direccion string) {
	var nuevaCapacidad int
	if direccion == "arriba" {
		nuevaCapacidad = p.cantidad * VALOR_PARA_AGRANDAR
	} else if direccion == "abajo" {
		nuevaCapacidad = len(p.datos) / VALOR_PARA_ACHICAR
	}
	if nuevaCapacidad == 0 {
		nuevaCapacidad = 1
	}
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, p.datos)
	p.datos = nuevosDatos
}
