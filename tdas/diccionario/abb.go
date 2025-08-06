package diccionario

import (
	TDAPila "tdas/pila"
)

/* Implementacion diccionario ABB */

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{nil, 0, funcion_cmp}
}

// crearNodoABB crea un nodo del arbol binario de busqueda con la clave y el dato indicados por parámetro.
func crearNodoABB[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{nil, nil, clave, dato}
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	nodo := buscarNodo(clave, &abb.raiz, abb.cmp)
	if *nodo == nil {
		*nodo = crearNodoABB(clave, dato)
		abb.cantidad++
	} else {
		(*nodo).dato = dato
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	return *buscarNodo(clave, &abb.raiz, abb.cmp) != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo := buscarNodo(clave, &abb.raiz, abb.cmp)
	panicAbb(nodo)
	return (*nodo).dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	borrado := buscarNodo(clave, &abb.raiz, abb.cmp)
	panicAbb(borrado)
	dato := (*borrado).dato
	borrar(borrado)
	abb.cantidad--

	return dato
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	iterar(abb.raiz, visitar, nil, nil, abb.cmp)
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	iterar(abb.raiz, visitar, desde, hasta, abb.cmp)
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	return crearIter(abb, desde, hasta)
}

// borrar borra un nodo del arbol binario de busqueda y analiza los casos posibles para mantener la propiedad de arbol binario de busqueda.
func borrar[K comparable, V any](nodo **nodoAbb[K, V]) {
	if (*nodo).izquierdo == nil {
		*nodo = (*nodo).derecho
		return
	}
	if (*nodo).derecho == nil {
		*nodo = (*nodo).izquierdo
		return
	}

	reemplazo := &(*nodo).derecho
	for (*reemplazo).izquierdo != nil {
		reemplazo = &(*reemplazo).izquierdo
	}

	(*nodo).clave = (*reemplazo).clave
	(*nodo).dato = (*reemplazo).dato

	borrar(reemplazo)
}

// iterar recorre el arbol segun el método in-order y aplica la función visitar a cada nodo cuya clave se encuentre en el rango indicado.
// Si la función visitar devuelve false, se interrumpe la iteración. Si los límites de rango son nil, no se los tiene en cuenta
func iterar[K comparable, V any](nodo *nodoAbb[K, V], visitar func(clave K, dato V) bool, desde *K, hasta *K, funcion_cmp func(K, K) int) bool {
	if nodo == nil {
		return true
	}
	if desde == nil || funcion_cmp(nodo.clave, *desde) >= 0 {
		if !iterar(nodo.izquierdo, visitar, desde, hasta, funcion_cmp) {
			return false
		}
	}
	if (desde == nil || funcion_cmp(nodo.clave, *desde) >= 0) && (hasta == nil || funcion_cmp(nodo.clave, *hasta) <= 0) {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}
	if hasta == nil || funcion_cmp(nodo.clave, *hasta) <= 0 {
		if !iterar(nodo.derecho, visitar, desde, hasta, funcion_cmp) {
			return false
		}
	}
	return true
}

// buscarNodo devuelve un puntero a un puntero al nodo cuya clave coincide con la indicada por parámetro.
// En caso de que el nodo no pertenezca al árbol, devuelve un puntero a la dirección dónde este debería encontrarse.
func buscarNodo[K comparable, V any](clave K, actual **nodoAbb[K, V], funcion_cmp func(K, K) int) **nodoAbb[K, V] {
	if *actual == nil || funcion_cmp(clave, (*actual).clave) == 0 {
		return actual
	}
	if funcion_cmp(clave, (*actual).clave) < 0 {
		return buscarNodo(clave, &((*actual).izquierdo), funcion_cmp)
	}
	return buscarNodo(clave, &((*actual).derecho), funcion_cmp)
}

// panicAbb lanza un panic si el nodo indicado por parámetro es nulo
func panicAbb[K comparable, V any](nodo **nodoAbb[K, V]) {
	if *nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
}

/* Implementacion iterador externo*/

type iterABB[K comparable, V any] struct {
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
	cmp   func(K, K) int
}

// crearIter crea un iterador del arbol binario de busqueda con el rango indicado por parámetro.
// Si no se indica un rango, se crea un iterador que recorre todo el arbol.
func crearIter[K comparable, V any](abb *abb[K, V], desde *K, hasta *K) *iterABB[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter := iterABB[K, V]{pila, desde, hasta, abb.cmp}
	apilarIzqRango(&iter, abb.raiz)
	return &iter
}

func (iter *iterABB[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iterABB[K, V]) VerActual() (K, V) {
	panicIterAbb(iter)

	actual := iter.pila.VerTope()
	return actual.clave, actual.dato
}

func (iter *iterABB[K, V]) Siguiente() {
	panicIterAbb(iter)
	nodo := iter.pila.Desapilar()
	if nodo.derecho != nil {
		apilarIzqRango(iter, nodo.derecho)
	}
}

// apilarIzqRango apila todos los sucesivos hijos izquierdos de nodo siempre y cuando sean >= desde y <= hasta.
// Si los izquierdos son menores, considera derechos.
func apilarIzqRango[K comparable, V any](iter *iterABB[K, V], nodo *nodoAbb[K, V]) {
	if nodo == nil {
		return
	}
	if iter.desde != nil && iter.cmp(nodo.clave, *iter.desde) < 0 {
		apilarIzqRango(iter, nodo.derecho)
		return
	}
	if iter.hasta != nil && iter.cmp(nodo.clave, *iter.hasta) > 0 {
		apilarIzqRango(iter, nodo.izquierdo)
		return
	}
	iter.pila.Apilar(nodo)
	apilarIzqRango(iter, nodo.izquierdo)
}

// panicIterAbb lanza un panic si el iterador terminó de iterar
func panicIterAbb[K comparable, V any](iter *iterABB[K, V]) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}
