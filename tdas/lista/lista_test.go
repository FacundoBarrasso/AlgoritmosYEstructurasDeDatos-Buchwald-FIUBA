package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	CANTIDAD_VOLUMEN        = 10000
	CONDICION_CORTE_INTERNO = 10
)

/* Tests de lista */

func TestListaVacia(t *testing.T) {
	t.Log("Evaluamos el comportamiento de una lista vacía.")
	lista := TDALista.CrearListaEnlazada[int]()
	require.NotNil(t, lista, "Al crear una lista esta no vale nil")
	testListaVacia(t, lista)
}

func TestPocosElementos(t *testing.T) {
	var (
		enterosPar   = []int{1, 2, 3, 4}
		enterosImpar = []int{1, 2, 3}
		floatPar     = []float64{1.123456, 0.789456, 3.789645, 4.8652543}
		floatImpar   = []float64{1.123456, 0.789456, 3.789645}
		cadenasPar   = []string{"Alma de Diamante", "Almendra", "Bajo Belgrano", "Peperina"}
		cadenasImpar = []string{"Pulp Fiction", "Lalaland", "Taxi Driver"}
		boolPar      = []bool{true, true, true, false}
		boolImpar    = []bool{false, true, false}
	)

	testListaUnElemento(t, enterosPar[0], func(lista TDALista.Lista[int], a int) { lista.InsertarPrimero(a) }, "enteros")
	testListaUnElemento(t, enterosPar[0], func(lista TDALista.Lista[int], a int) { lista.InsertarUltimo(a) }, "enteros")

	testListaUnElemento(t, enterosImpar[0], func(lista TDALista.Lista[int], a int) { lista.InsertarPrimero(a) }, "enteros")
	testListaUnElemento(t, enterosImpar[0], func(lista TDALista.Lista[int], a int) { lista.InsertarUltimo(a) }, "enteros")

	testListaUnElemento(t, floatPar[0], func(lista TDALista.Lista[float64], a float64) { lista.InsertarPrimero(a) }, "floats")
	testListaUnElemento(t, floatPar[0], func(lista TDALista.Lista[float64], a float64) { lista.InsertarUltimo(a) }, "floats")

	testListaUnElemento(t, floatImpar[0], func(lista TDALista.Lista[float64], a float64) { lista.InsertarPrimero(a) }, "floats")
	testListaUnElemento(t, floatImpar[0], func(lista TDALista.Lista[float64], a float64) { lista.InsertarUltimo(a) }, "floats")

	testListaUnElemento(t, cadenasPar[0], func(lista TDALista.Lista[string], a string) { lista.InsertarPrimero(a) }, "cadenas")
	testListaUnElemento(t, cadenasPar[0], func(lista TDALista.Lista[string], a string) { lista.InsertarUltimo(a) }, "cadenas")

	testListaUnElemento(t, cadenasImpar[0], func(lista TDALista.Lista[string], a string) { lista.InsertarPrimero(a) }, "cadenas")
	testListaUnElemento(t, cadenasImpar[0], func(lista TDALista.Lista[string], a string) { lista.InsertarUltimo(a) }, "cadenas")

	testListaUnElemento(t, boolPar[0], func(lista TDALista.Lista[bool], a bool) { lista.InsertarPrimero(a) }, "booleanos")
	testListaUnElemento(t, boolPar[0], func(lista TDALista.Lista[bool], a bool) { lista.InsertarUltimo(a) }, "booleanos")

	testListaUnElemento(t, boolImpar[0], func(lista TDALista.Lista[bool], a bool) { lista.InsertarPrimero(a) }, "booleanos")
	testListaUnElemento(t, boolImpar[0], func(lista TDALista.Lista[bool], a bool) { lista.InsertarUltimo(a) }, "booleanos")

	testListaPocosElementos(t, enterosPar, func(lista TDALista.Lista[int], a int) { lista.InsertarPrimero(a) }, func(lista TDALista.Lista[int]) int { return lista.VerPrimero() }, "enteros")
	testListaPocosElementos(t, enterosPar, func(lista TDALista.Lista[int], a int) { lista.InsertarUltimo(a) }, func(lista TDALista.Lista[int]) int { return lista.VerUltimo() }, "enteros")

	testListaPocosElementos(t, enterosImpar, func(lista TDALista.Lista[int], a int) { lista.InsertarPrimero(a) }, func(lista TDALista.Lista[int]) int { return lista.VerPrimero() }, "enteros")
	testListaPocosElementos(t, enterosImpar, func(lista TDALista.Lista[int], a int) { lista.InsertarUltimo(a) }, func(lista TDALista.Lista[int]) int { return lista.VerUltimo() }, "enteros")

	testListaPocosElementos(t, floatPar, func(lista TDALista.Lista[float64], a float64) { lista.InsertarPrimero(a) }, func(lista TDALista.Lista[float64]) float64 { return lista.VerPrimero() }, "floats")
	testListaPocosElementos(t, floatPar, func(lista TDALista.Lista[float64], a float64) { lista.InsertarUltimo(a) }, func(lista TDALista.Lista[float64]) float64 { return lista.VerUltimo() }, "floats")

	testListaPocosElementos(t, floatImpar, func(lista TDALista.Lista[float64], a float64) { lista.InsertarPrimero(a) }, func(lista TDALista.Lista[float64]) float64 { return lista.VerPrimero() }, "floats")
	testListaPocosElementos(t, floatImpar, func(lista TDALista.Lista[float64], a float64) { lista.InsertarUltimo(a) }, func(lista TDALista.Lista[float64]) float64 { return lista.VerUltimo() }, "floats")

	testListaPocosElementos(t, cadenasPar, func(lista TDALista.Lista[string], a string) { lista.InsertarPrimero(a) }, func(lista TDALista.Lista[string]) string { return lista.VerPrimero() }, "cadenas")
	testListaPocosElementos(t, cadenasPar, func(lista TDALista.Lista[string], a string) { lista.InsertarUltimo(a) }, func(lista TDALista.Lista[string]) string { return lista.VerUltimo() }, "cadenas")

	testListaPocosElementos(t, cadenasImpar, func(lista TDALista.Lista[string], a string) { lista.InsertarPrimero(a) }, func(lista TDALista.Lista[string]) string { return lista.VerPrimero() }, "cadenas")
	testListaPocosElementos(t, cadenasImpar, func(lista TDALista.Lista[string], a string) { lista.InsertarUltimo(a) }, func(lista TDALista.Lista[string]) string { return lista.VerUltimo() }, "cadenas")

	testListaPocosElementos(t, boolPar, func(lista TDALista.Lista[bool], a bool) { lista.InsertarPrimero(a) }, func(lista TDALista.Lista[bool]) bool { return lista.VerPrimero() }, "booleanos")
	testListaPocosElementos(t, boolPar, func(lista TDALista.Lista[bool], a bool) { lista.InsertarUltimo(a) }, func(lista TDALista.Lista[bool]) bool { return lista.VerUltimo() }, "booleanos")

	testListaPocosElementos(t, boolImpar, func(lista TDALista.Lista[bool], a bool) { lista.InsertarPrimero(a) }, func(lista TDALista.Lista[bool]) bool { return lista.VerPrimero() }, "booleanos")
	testListaPocosElementos(t, boolImpar, func(lista TDALista.Lista[bool], a bool) { lista.InsertarUltimo(a) }, func(lista TDALista.Lista[bool]) bool { return lista.VerUltimo() }, "booleanos")
}

func TestListaVolumen(t *testing.T) {
	t.Logf("Evaluamos el comportamiento de una lista con %d elementos.", CANTIDAD_VOLUMEN)
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i < CANTIDAD_VOLUMEN; i++ {
		lista.InsertarUltimo(i)
		require.Equal(t, i, lista.VerUltimo(), "El resultado de insertar un elemento al final de la lista coincide"+
			"con ver el elemento último elemento")
		require.Equal(t, 0, lista.VerPrimero(), "Al insertar elementos al final el primer elemento permanece igual")
	}
	require.Equal(t, CANTIDAD_VOLUMEN, lista.Largo(), "El largo de la lista debe coincidir con la cantidad de elementos insertados")

	contador := 0
	lista.Iterar(func(n int) bool {
		require.Equal(t, contador, n, "El elemento iterado debe coincidir con el valor esperado")
		contador++
		return true
	})
	require.Equal(t, CANTIDAD_VOLUMEN, contador, "Se deben haber iterado todos los elementos")

	for i := 0; i < CANTIDAD_VOLUMEN; i++ {
		require.Equal(t, i, lista.BorrarPrimero(), "El elemento borrado coincide con el valor esperado")
	}

	testListaVacia(t, lista)
}

func TestIteradorInternoSinCorte(t *testing.T) {
	var (
		elementos           = []int{1, 2, 3, 4, 5}
		resultadosEsperados = []int{1, 3, 6, 10, 15}
		resultadosParciales []int
	)

	t.Log("Evaluamos el iterador interno sin condición de corte.")

	lista := TDALista.CrearListaEnlazada[int]()
	for _, elem := range elementos {
		lista.InsertarUltimo(elem)
	}

	contador := 0
	lista.Iterar(func(n int) bool {
		contador += n
		resultadosParciales = append(resultadosParciales, contador)
		return true
	})

	require.Equal(t, resultadosEsperados, resultadosParciales, "Los resultados parciales coinciden con lo esperado en cada iteración")
}

func TestIteradorInternoConCorte(t *testing.T) {
	var (
		elementos           = []int{1, 2, 3, 4, 5, 6}
		resultadosEsperados = []int{1, 3, 6, 10}
		resultadosParciales []int
	)

	t.Log("Evaluamos el iterador interno con condición de corte.")

	lista := TDALista.CrearListaEnlazada[int]()
	for _, elem := range elementos {
		lista.InsertarUltimo(elem)
	}

	contador := 0
	lista.Iterar(func(n int) bool {
		contador += n
		resultadosParciales = append(resultadosParciales, contador)
		return contador < CONDICION_CORTE_INTERNO
	})

	require.Equal(t, resultadosEsperados, resultadosParciales, "Los resultados parciales coinciden con lo esperado en cada iteración")
}

func testListaVacia[T any](t *testing.T, lista TDALista.Lista[T]) {
	require.True(t, lista.EstaVacia(), "Si la lista no tiene ningun elemento, entonces está vacía")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "Si la "+
		"lista está vacía, entra en pánico al borrar el primer elemento")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "Si la "+
		"lista está vacía, entra en pánico al ver el primer elemento")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "Si la "+
		"lista está vacía, entra en pánico al ver el último elemento")
	require.Equal(t, 0, lista.Largo(), "El largo de una lista vacía es cero")

	var contador int
	lista.Iterar(func(n T) bool { contador++; return true })
	require.Equal(t, 0, contador, "Resulta trivial iterar sobre una lista vacía")
}

func testListaUnElemento[T any](t *testing.T, a T, insertar func(lista TDALista.Lista[T], a T), tipoDato string) {
	t.Logf("Evaluamos el comportamiento de una lista de %s con un único elemento", tipoDato)
	lista := TDALista.CrearListaEnlazada[T]()
	insertar(lista, a)
	require.False(t, lista.EstaVacia(), "Una lista con un elemento no está vacía")
	require.EqualValues(t, a, lista.VerPrimero(), "Al insertar un nuevo dato en una lista vacia este es el primer elemento")
	require.False(t, lista.EstaVacia(), "No se elimina el elemento de la lista luego de ver el primero")
	require.EqualValues(t, a, lista.VerUltimo(), "Al insertar un nuevo dato en una lista vacia este es el último elemento")
	require.False(t, lista.EstaVacia(), "No se elimina el elemento de la lista luego de ver el último")
	require.Equal(t, 1, lista.Largo(), "El largo de una lista con un único elemento es uno")

	arreglo := make([]T, 0, 1)
	lista.Iterar(func(n T) bool { arreglo = append(arreglo, n); return true })
	require.Equal(t, a, arreglo[0], "El valor del primer elemento de un arreglo coincide con el único elemento "+
		"de la lista, tras iterarla y almacenarla en el arreglo")
	require.Equal(t, 1, len(arreglo), "El largo de un arreglo es uno, tras iterar una lista de un"+
		"elemento y almacenarla en el arreglo")

	require.Equal(t, a, lista.BorrarPrimero(), "Se devuelve el único elemento de la lista al borrarlo")
	testListaVacia(t, lista)
}

func testListaPocosElementos[T any](t *testing.T, elementos []T, insertar func(lista TDALista.Lista[T], a T), ver func(lista TDALista.Lista[T]) T, tipoDato string) {
	t.Logf("Evaluamos el comportamiento de una lista de %s con %d elementos", tipoDato, len(elementos))
	lista := TDALista.CrearListaEnlazada[T]()

	for i, elem := range elementos {
		insertar(lista, elem)
		require.Equal(t, elementos[i], ver(lista), "El resultado de insertar un elemento en una posición coincide"+
			"con ver el elemento de esa misma posición")
	}
	require.Equal(t, len(elementos), lista.Largo(), "El largo de la lista coincide con la cantidad de elementos insertados")

	elemsInvertido := invertirArreglo(elementos)
	arreglo := make([]T, 0, len(elementos))
	lista.Iterar(func(n T) bool { arreglo = append(arreglo, n); return true })
	require.Contains(t, [][]T{elementos, elemsInvertido}, arreglo, "Los elementos iterados coinciden con los insertados")

	for i := 0; i < len(elementos); i++ {
		require.Contains(t, []T{elementos[i], elemsInvertido[i]}, lista.BorrarPrimero(), "El elemento borrado debe coincide con el primer elemento insertado")
	}

	testListaVacia(t, lista)
}

func invertirArreglo[T any](arreglo []T) []T {
	n := len(arreglo)
	invertido := make([]T, n)
	for i := 0; i < n; i++ {
		invertido[i] = arreglo[n-i-1]
	}
	return invertido
}

/* Tests iterador externo */

func TestIteradorInsertarAlPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	primero := 28
	lista.InsertarUltimo(primero)
	lista.InsertarUltimo(39)
	lista.InsertarUltimo(13)

	t.Log("Al insertar un elemento en la posición en la que se crea el iterador, debe insertarse al principio.")

	insertado := 100
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if iter.VerActual() == primero {
			iter.Insertar(insertado)
			break
		}
	}
	require.Equal(t, insertado, lista.VerPrimero(), "El elemento insertado en la primer posicion del iterador "+
		"es el primero de la lista.")
}

func TestIteradorInsertarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	primero := 20
	lista.InsertarUltimo(primero)

	t.Log("Insertar un elemento cuando el iterador está al final debe ser equivalente a insertar al final.")

	insertado := 13
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	iter.Insertar(insertado)
	require.Equal(t, primero, lista.VerPrimero(), "El primer elemento de la lista es el insertado inicialmente.")
	require.Equal(t, insertado, lista.VerUltimo(), "El elemento insertado en la ultima posicion del iterador "+
		"es el ultimo de la lista.")
}

func TestIteradorInsertarEnMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	primero := 10
	lista.InsertarUltimo(primero)
	ultimo := 39
	lista.InsertarUltimo(ultimo)

	t.Log("Insertar un elemento en el medio de la iteración, se hace en la posición correcta.")

	insertado := 26
	cont := 0
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if cont == 1 {
			iter.Insertar(insertado)
			break
		}
		cont++
	}
	require.Equal(t, primero, lista.VerPrimero(), "El primer elemento de la lista es el insertado inicialmente.")
	require.Equal(t, ultimo, lista.VerUltimo(), "El último elemento de la lista es el insertado inicialmente.")

	cont = 0
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if cont == 1 {
			require.Equal(t, insertado, iter.VerActual(), "El elemento insertado está en la posicion 1 de la lista.")
			break
		}
		cont++
	}
}

func TestIteradorBorrarAlPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	primero := 6
	lista.InsertarUltimo(primero)
	segundo := 28
	lista.InsertarUltimo(segundo)
	lista.InsertarUltimo(789)

	t.Log("Al borrar el elemento cuando se crea el iterador, cambia el primer elemento de la lista.")

	cont := 0
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if cont == 0 {
			borrado := iter.Borrar()
			require.Equal(t, primero, borrado, "El elemento borrado es el primero insertado en la lista.")
		}
		cont++
	}
	require.Equal(t, segundo, lista.VerPrimero(), "El primer elemento de la lista cambia al borrar el primer "+
		"elemento de la iteración.")
}

func TestIteradorBorrarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(10)
	anteultimo := 82
	lista.InsertarUltimo(anteultimo)
	ultimo := 38
	lista.InsertarUltimo(ultimo)

	t.Log("Borrar el último elemento con el iterador, cambia el último elemento de la lista.")

	cont := 0
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if cont == lista.Largo()-1 {
			require.Equal(t, ultimo, lista.VerUltimo(), "El último elemento de la lista es el ultimo insertado.")
			borrado := iter.Borrar()
			require.Equal(t, ultimo, borrado, "El elemento borrado es el ultimo insertado en la lista.")
			require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() }, "El elemento actual paso "+
				"a ser nil al eliminar el último")
			break
		}
		cont++
	}
	require.Equal(t, anteultimo, lista.VerUltimo(), "El último elemento de la lista cambia al borrar el ultimo "+
		"elemento de la iteración.")

	cont = 0
	iter := lista.Iterador()
	for cont < lista.Largo() {
		iter.Borrar()
		cont++
	}
	require.Equal(t, anteultimo, lista.VerUltimo(), "Al borrar todos los elementos menos el anteúltimo de la lista "+
		"original, este pasa a ser el último.")
	require.Equal(t, anteultimo, lista.VerPrimero(), "Al borrar todos los elementos menos el anteúltimo de la lista "+
		"original, este pasa a ser el primero.")
	require.Equal(t, anteultimo, iter.Borrar(), "Se borra el elemento que queda en la lista.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() }, "El elemento actual paso "+
		"a ser nil al eliminar el último")
	testListaVacia(t, lista)

}

func TestIteradorBorrarEnMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(11)
	medio := 22
	lista.InsertarUltimo(medio)
	lista.InsertarUltimo(32)
	lista.InsertarUltimo(23)

	t.Log("Al borrar un elemento del medio, este no está.")

	cont := 0
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if cont == 1 {
			borrado := iter.Borrar()
			require.Equal(t, medio, borrado, "El elemento borrado es el del medio de la lista.")
			require.NotEqual(t, medio, iter.VerActual(), "El elemento borrado no está en la lista.")
			break
		}
		cont++
	}
}

func TestIteradorListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	t.Log("Verificar que el iterador de una lista vacía no tiene elementos.")

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	require.False(t, iter.HaySiguiente(), "El iterador de una lista vacía no tiene elementos.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.Siguiente()
	}, "No se puede avanzar en un iterador de una lista vacía.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.Borrar()
	}, "No se puede borrar un elemento en un iterador de una lista vacía.")
}

func TestIteradorTerminoDeIterar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)

	t.Log("Verificar que el iterador termina de iterar.")

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.Siguiente()
	}, "El iterador termina de iterar al llegar al final de la lista.")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.Borrar()
	}, "El iterador termina de iterar al llegar al final de la lista.")
}

func TestIteradorInsertarEnListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	t.Log("Insertar un elemento en una lista vacía, el elemento insertado pasa a ser el primero y el ultimo.")

	iter := lista.Iterador()
	insertado := 10
	iter.Insertar(insertado)
	require.Equal(t, insertado, lista.VerPrimero(), "El elemento insertado es el primero de la lista.")
	require.Equal(t, insertado, lista.VerUltimo(), "El elemento insertado es el ultimo de la lista.")
}

func TestIteradorBorrarEnListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	t.Log("Borrar un elemento en una lista vacía, no hace nada.")

	iter := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.Borrar()
	}, "No se puede borrar un elemento en una lista vacía.")
}

func TestIteradorBorrarUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	unico := 10
	lista.InsertarUltimo(unico)

	t.Log("Borrar un elemento en una lista con un único elemento, la lista queda vacía y permite su uso normal.")

	iter := lista.Iterador()

	var borrado int
	for iter.HaySiguiente() {
		borrado = iter.Borrar()
	}

	require.Equal(t, unico, borrado, "El elemento borrado es el único de la lista.")
	testListaVacia(t, lista)

	lista.InsertarUltimo(2890)
	require.Equal(t, 2890, lista.VerPrimero(), "La primitiva de VerPrimero funciona luego de vaciar la lista y volver a insertar.")
	require.Equal(t, 2890, lista.VerUltimo(), "La primitiva de VerUltimo funciona luego de vaciar la lista y volver a insertar.")
}
