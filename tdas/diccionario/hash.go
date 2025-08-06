package diccionario

import (
	"fmt"
	"hash/fnv"
)

type estadoParClaveDato int

const (
	VACIO estadoParClaveDato = iota
	BORRADO
	OCUPADO
	TAMANIO_INICIAL       = 7
	FACTOR_CARGA_SUPERIOR = 0.7
	FACTOR_CARGA_INFERIOR = 0.3
	FACTOR_REDIMENSION    = 2
	VALOR_INICIAL         = 0
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estadoParClaveDato
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	borrados int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashCerrado[K, V]{crearTabla[K, V](TAMANIO_INICIAL), VALOR_INICIAL, VALOR_INICIAL}
}

func crearTabla[K comparable, V any](capacidad int) []celdaHash[K, V] {
	return make([]celdaHash[K, V], capacidad)
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	if float64(hash.cantidad+hash.borrados)/float64(len(hash.tabla)) >= FACTOR_CARGA_SUPERIOR {
		hash.redimensionar(len(hash.tabla) * FACTOR_REDIMENSION)
	}
	posicion := buscarSiguientePosicion(hash.tabla, clave)
	if hash.tabla[posicion].estado == VACIO {
		hash.tabla[posicion].clave = clave
		hash.tabla[posicion].estado = OCUPADO
		hash.cantidad++
	}
	hash.tabla[posicion].dato = dato
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	posicion := buscarSiguientePosicion(hash.tabla, clave)
	return hash.tabla[posicion].estado != VACIO
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	posicion := buscarSiguientePosicion(hash.tabla, clave)
	panicPertenece(hash.tabla[posicion].estado)
	return hash.tabla[posicion].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	if float32(hash.cantidad)/float32(len(hash.tabla)) <= FACTOR_CARGA_INFERIOR && len(hash.tabla) > TAMANIO_INICIAL {
		hash.redimensionar(len(hash.tabla) / FACTOR_REDIMENSION)
	}

	posicion := buscarSiguientePosicion(hash.tabla, clave)
	panicPertenece(hash.tabla[posicion].estado)
	hash.tabla[posicion].estado = BORRADO
	hash.cantidad--
	hash.borrados++

	return hash.tabla[posicion].dato
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for _, celda := range hash.tabla {
		if celda.estado == OCUPADO {
			if !visitar(celda.clave, celda.dato) {
				break
			}
		}
	}
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iterHash[K, V]{hash, VALOR_INICIAL}
	iter.avanzarSigOcupado()
	return iter
}

// redimensionar ajusta el tamaño de la tabla hash a un nuevo tamaño especificado, limpiando los elementos borrados.
func (hash *hashCerrado[K, V]) redimensionar(nuevoTam int) {
	tablaNueva := crearTabla[K, V](nuevoTam)
	for _, celda := range hash.tabla {
		if celda.estado == OCUPADO {
			posicion := buscarSiguientePosicion(tablaNueva, celda.clave)
			tablaNueva[posicion].clave = celda.clave
			tablaNueva[posicion].dato = celda.dato
			tablaNueva[posicion].estado = OCUPADO
		}
	}
	hash.tabla = tablaNueva
	hash.borrados = VALOR_INICIAL
}

// buscarSiguientePosicion busca la siguiente posición vacía o la posición de una clave específica en una tabla hash.
func buscarSiguientePosicion[K comparable, V any](tabla []celdaHash[K, V], clave K) int {
	pos := int(hashing(clave, len(tabla)))
	vuelta := false
	for i := pos; i < pos || !vuelta; i++ {
		if (tabla[i].estado == OCUPADO && tabla[i].clave == clave) || tabla[i].estado == VACIO {
			return i
		} else if i+1 == len(tabla) {
			i = 0
			vuelta = true
		}
	}
	return -1
}

// panicPertenece lanza un pánico si la clave no pertenece al diccionario.
func panicPertenece(estado estadoParClaveDato) {
	if estado == VACIO {
		panic("La clave no pertenece al diccionario")
	}
}

/* Definición estructura iterador externo */

type iterHash[K comparable, V any] struct {
	hash   *hashCerrado[K, V]
	actual int
}

func (iter *iterHash[K, V]) HaySiguiente() bool {
	return iter.actual < len(iter.hash.tabla)
}

func (iter *iterHash[K, V]) VerActual() (K, V) {
	panicIter(iter)
	return iter.hash.tabla[iter.actual].clave, iter.hash.tabla[iter.actual].dato
}

func (iter *iterHash[K, V]) Siguiente() {
	panicIter(iter)
	iter.actual++
	iter.avanzarSigOcupado()
}

// avanzarSigOcupado si existe una proxima celda ocupada avanza el iterador a esa posición,
// en caso contrario deja al actual fuera de rango del iterador.
func (iter *iterHash[K, V]) avanzarSigOcupado() {
	for i := iter.actual; i <= len(iter.hash.tabla); i++ {
		if i == len(iter.hash.tabla) || iter.hash.tabla[i].estado == OCUPADO {
			iter.actual = i
			break
		}
	}
}

// panicIter lanza un pánico si el iterador terminó de iterar.
func panicIter[K comparable, V any](iter *iterHash[K, V]) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

/* Funciones para hashing */

// hashing devuelve el hash de una clave en un rango específico.
func hashing[K comparable](clave K, tam int) uint32 {
	str := fmt.Sprintf("%v", clave)
	return fnv32(str) % uint32(tam)
}

// Magic
func fnv32(str string) uint32 {
	hash := fnv.New32()
	hash.Write([]byte(str))
	return hash.Sum32()
}
