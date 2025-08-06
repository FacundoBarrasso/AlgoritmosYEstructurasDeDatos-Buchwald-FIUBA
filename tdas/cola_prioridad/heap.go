package cola_prioridad

const (
	FACTOR_REDIMENSION = 2
	FACTOR_REDUCCION   = 4
	TAMANIO_INICIAL    = 7
)

type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

func CrearHeap[T any](cmp func(T, T) int) ColaPrioridad[T] {
	return CrearHeapArr([]T{}, cmp)
}

func CrearHeapArr[T any](arreglo []T, cmp func(T, T) int) ColaPrioridad[T] {
	tamanio := TAMANIO_INICIAL
	if len(arreglo) > 0 {
		tamanio = FACTOR_REDIMENSION * len(arreglo)
	}
	datos := make([]T, tamanio)
	copy(datos, arreglo)
	heapify(&datos, len(arreglo), cmp)
	return &heap[T]{datos, len(arreglo), cmp}
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) Encolar(elem T) {
	if heap.cantidad == cap(heap.datos) {
		heap.redimensionar(cap(heap.datos) * FACTOR_REDIMENSION)
	}
	heap.datos[heap.cantidad] = elem
	upheap(heap, heap.cantidad)
	heap.cantidad++
}

func (heap *heap[T]) VerMax() T {
	heap.panicHeap()
	return heap.datos[0]
}

func (heap *heap[T]) Desencolar() T {
	heap.panicHeap()
	heap.cantidad--
	if cap(heap.datos) >= heap.cantidad*FACTOR_REDUCCION {
		heap.redimensionar(cap(heap.datos) / FACTOR_REDIMENSION)
	}
	elemento := heap.datos[0]
	heap.datos[0], heap.datos[heap.cantidad] = heap.datos[heap.cantidad], heap.datos[0]
	downheap(0, &heap.datos, heap.cantidad, heap.cmp)
	return elemento
}

func (heap *heap[T]) Cantidad() int {
	return heap.cantidad
}

func downheap[T any](posicion int, arr *[]T, cantidad int, cmp func(T, T) int) {
	maximo := posicion
	if posicion*2+1 < cantidad {
		maximo = maxPosicion(maximo, posicion*2+1, arr, cmp)
	}
	if posicion*2+2 < cantidad {
		maximo = maxPosicion(maximo, posicion*2+2, arr, cmp)
	}
	if maximo != posicion {
		(*arr)[posicion], (*arr)[maximo] = (*arr)[maximo], (*arr)[posicion]
		downheap[T](maximo, arr, cantidad, cmp)
	}
}

func maxPosicion[T any](a, b int, arr *[]T, cmp func(T, T) int) int {
	if cmp((*arr)[a], (*arr)[b]) >= 0 {
		return a
	}
	return b
}

func heapify[T any](arr *[]T, cantidad int, cmp func(T, T) int) {
	for i := (cantidad - 1) / 2; i >= 0; i-- {
		downheap(i, arr, cantidad, cmp)
	}
}

func upheap[T any](heap *heap[T], i int) {
	for i > 0 {
		padre := (i - 1) / 2
		if heap.cmp(heap.datos[padre], heap.datos[i]) < 0 {
			swap(heap.datos, padre, i)
			i = padre
		} else {
			break
		}
	}
}

func HeapSort[T any](arreglo []T, cmp func(T, T) int) {
	heapify(&arreglo, len(arreglo), cmp)
	for i := len(arreglo) - 1; i > 0; i-- {
		swap(arreglo, 0, i)
		downheap(0, &arreglo, i, cmp)
	}
}

func (heap *heap[T]) panicHeap() {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
}

func (heap *heap[T]) redimensionar(tamanio int) {
	if tamanio < TAMANIO_INICIAL {
		tamanio = TAMANIO_INICIAL
	}
	nuevosDatos := make([]T, tamanio)
	copy(nuevosDatos, heap.datos)
	heap.datos = nuevosDatos
}

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
