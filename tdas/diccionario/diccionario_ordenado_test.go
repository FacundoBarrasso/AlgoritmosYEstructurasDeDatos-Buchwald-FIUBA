package diccionario_test

import (
	"fmt"
	"math/rand"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"slices"

	"github.com/stretchr/testify/require"
)

/* Tests de ABB */

var TAMANIOS_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

func cmpInt(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func mezclarClaves[V any](claves []V) []V {
	n := len(claves)
	for i := 0; i < n; i++ {
		j := rand.Intn(n)
		claves[i], claves[j] = claves[j], claves[i]
	}
	return claves
}

func TestABBVacio(t *testing.T) {
	t.Log("Comprueba que un ABB vacio no tiene claves")
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("A") })
}

func TestABBClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un ABB vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, abb.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("") })

	abbNum := TDADiccionario.CrearABB[int, string](cmpInt)
	require.False(t, abbNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbNum.Borrar(0) })
}

func TestABBUnElemento(t *testing.T) {
	t.Log("Comprueba que ABB con un elemento tiene esa Clave, unicamente")
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	abb.Guardar("A", 10)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece("A"))
	require.False(t, abb.Pertenece("B"))
	require.EqualValues(t, 10, abb.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("B") })
}

func TestABBGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el ABB, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, abb.Pertenece(claves[0]))
	abb.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))

	require.False(t, abb.Pertenece(claves[1]))
	require.False(t, abb.Pertenece(claves[2]))
	abb.Guardar(claves[1], valores[1])
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], abb.Obtener(claves[1]))

	require.False(t, abb.Pertenece(claves[2]))
	abb.Guardar(claves[2], valores[2])
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, 3, abb.Cantidad())
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], abb.Obtener(claves[1]))
	require.EqualValues(t, valores[2], abb.Obtener(claves[2]))
}

func TestABBReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	abb.Guardar(clave, "miau")
	abb.Guardar(clave2, "guau")
	require.True(t, abb.Pertenece(clave))
	require.True(t, abb.Pertenece(clave2))
	require.EqualValues(t, "miau", abb.Obtener(clave))
	require.EqualValues(t, "guau", abb.Obtener(clave2))
	require.EqualValues(t, 2, abb.Cantidad())

	abb.Guardar(clave, "miu")
	abb.Guardar(clave2, "baubau")
	require.True(t, abb.Pertenece(clave))
	require.True(t, abb.Pertenece(clave2))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, "miu", abb.Obtener(clave))
	require.EqualValues(t, "baubau", abb.Obtener(clave2))
}

func TestABBReemplazoDatoHopscotch(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	abb := TDADiccionario.CrearABB[int, int](cmpInt)

	claves := make([]int, 500)
	mezclarClaves(claves)

	for i := 0; i < 500; i++ {
		abb.Guardar(claves[i], claves[i])
	}
	for i := 0; i < 500; i++ {
		abb.Guardar(claves[i], 2*claves[i])
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = abb.Obtener(claves[i]) == 2*claves[i]
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestABBBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el ABB, y se los borra, revisando que en todo momento " +
		"el ABB se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)

	require.False(t, abb.Pertenece(claves[0]))
	require.False(t, abb.Pertenece(claves[0]))
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])

	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], abb.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[2]) })
	require.EqualValues(t, 2, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[2]))

	require.True(t, abb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], abb.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[0]) })
	require.EqualValues(t, 1, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claves[0]) })

	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], abb.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[1]) })
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claves[1]) })
}

func TestABBConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	clave := 10
	valor := "Gatito"

	abb.Guardar(clave, valor)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, valor, abb.Obtener(clave))
	require.EqualValues(t, valor, abb.Borrar(clave))
	require.False(t, abb.Pertenece(clave))
}

func TestABBConClavesStructs(t *testing.T) {
	t.Log("Valida que tambien funcione con estructuras mas complejas")
	type basico struct {
		a string
		b int
	}
	type avanzado struct {
		w int
		x basico
		y basico
		z string
	}

	abb := TDADiccionario.CrearABB[avanzado, int](func(a, b avanzado) int {
		if a.w < b.w {
			return -1
		} else if a.w > b.w {
			return 1
		}
		return strings.Compare(a.z, b.z)
	})

	a1 := avanzado{w: 10, z: "hola", x: basico{a: "mundo", b: 8}, y: basico{a: "!", b: 10}}
	a2 := avanzado{w: 10, z: "aloh", x: basico{a: "odnum", b: 14}, y: basico{a: "!", b: 5}}
	a3 := avanzado{w: 10, z: "hello", x: basico{a: "world", b: 8}, y: basico{a: "!", b: 4}}

	abb.Guardar(a1, 0)
	abb.Guardar(a2, 1)
	abb.Guardar(a3, 2)

	require.True(t, abb.Pertenece(a1))
	require.True(t, abb.Pertenece(a2))
	require.True(t, abb.Pertenece(a3))
	require.EqualValues(t, 0, abb.Obtener(a1))
	require.EqualValues(t, 1, abb.Obtener(a2))
	require.EqualValues(t, 2, abb.Obtener(a3))
	abb.Guardar(a1, 5)
	require.EqualValues(t, 5, abb.Obtener(a1))
	require.EqualValues(t, 2, abb.Obtener(a3))
	require.EqualValues(t, 5, abb.Borrar(a1))
	require.False(t, abb.Pertenece(a1))
	require.EqualValues(t, 2, abb.Obtener(a3))
}

func TestABBClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := ""
	abb.Guardar(clave, clave)
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, clave, abb.Obtener(clave))
}

func TestABBValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	abb := TDADiccionario.CrearABB[string, *int](strings.Compare)
	clave := "Pez"
	abb.Guardar(clave, nil)
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, (*int)(nil), abb.Obtener(clave))
	require.EqualValues(t, (*int)(nil), abb.Borrar(clave))
	require.False(t, abb.Pertenece(clave))
}

func ejecutarPruebaVolumenABB(b *testing.B, n int) {
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)

	claves := make([]string, n)
	valores := make([]int, n)

	/* Guarda "n" parejas en los arreglos */
	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = fmt.Sprintf("%08d", i)
	}

	mezclarClaves(claves) // Mezcla las claves para insertarlas en un orden aleatorio

	/* Inserta "n" parejas en el diccionario */
	for i := 0; i < n; i++ {
		abb.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, abb.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = abb.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = abb.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, abb.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = abb.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
		ok = !abb.Pertenece(claves[i])
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, abb.Cantidad())
}

func BenchmarkABB(b *testing.B) {
	b.Log("Prueba de stress del ABB. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves generadas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMANIOS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenABB(b, n)
			}
		})
	}
}

/* Tests de iterador interno */

func TestIteradorInternoVacio(t *testing.T) {
	t.Log("Valida que el iterador interno no haga nada si el ABB está vacío")
	abb := TDADiccionario.CrearABB[int, int](cmpInt)

	suma := 0
	abb.Iterar(func(_ int, dato int) bool {
		suma += dato
		return true
	})

	require.Equal(t, 0, suma)
}

func TestIteradorInternoSuma(t *testing.T) {
	t.Log("Valida que el iterador interno pueda sumar todos los elementos del ABB")
	abb := TDADiccionario.CrearABB[int, int](cmpInt)

	claves := []int{5, 3, 8, 1, 6, 4, 7, 2, 9, 10}

	for _, clave := range claves {
		abb.Guardar(clave, clave)
	}

	suma := 0
	abb.Iterar(func(_ int, dato int) bool {
		suma += dato
		return true
	})

	require.EqualValues(t, 55, suma)
}

func TestIteradorInternoCorte(t *testing.T) {
	t.Log("Valida que el iterador interno pueda cortar la iteración cuando se cumple una condición")
	abb := TDADiccionario.CrearABB[int, int](cmpInt)

	claves := []int{5, 3, 8, 1, 6, 4, 7, 2, 9, 10}

	for _, clave := range claves {
		abb.Guardar(clave, clave)
	}

	suma := 0
	cont := 0
	abb.Iterar(func(_ int, dato int) bool {
		suma += dato
		cont++
		return suma < 15 // Cortar cuando la suma sea mayor o igual a 15, en este caso debe frenar con 1+2+3+4+5 = 15
	})
	require.Equal(t, 5, cont) // Debe haber iterado 5 elementos
	require.Equal(t, 15, suma)
}

func TestIteradorInternoSinRangos(t *testing.T) {
	t.Log("Valida que el iterador interno funcione sin rangos y pueda multiplicar todos los elementos del ABB")
	abb := TDADiccionario.CrearABB[int, int](cmpInt)

	claves := []int{5, 3, 8, 1, 6, 4, 7, 2, 9, 10}

	for _, clave := range claves {
		abb.Guardar(clave, clave)
	}

	producto := 1
	i := 0
	abb.IterarRango(nil, nil, func(clave int, dato int) bool {
		producto *= dato
		require.Equal(t, clave, i+1) // Check that the keys are in ascending order
		i++
		return true
	})

	require.EqualValues(t, 3628800, producto) // 10! = 3628800
}

func TestIteradorInternoConRangos(t *testing.T) {
	t.Log("Valida que el iterador interno funcione con rangos")
	abb := TDADiccionario.CrearABB[int, int](cmpInt)

	claves := []int{5, 3, 8, 1, 6, 4, 7, 2, 9, 10}

	for _, clave := range claves {
		abb.Guardar(clave, clave)
	}

	desde := 3
	hasta := 7
	suma := 0
	abb.IterarRango(&desde, &hasta, func(_ int, dato int) bool {
		suma += dato
		return true
	})

	require.EqualValues(t, 25, suma) // 3 + 4 + 5 + 6 + 7 = 25
}

func TestIteradorInternoRangoConCorte(t *testing.T) {
	t.Log("Valida que el iterador interno pueda cortar la iteración dentro de un rango cuando se cumple una condición")
	abb := TDADiccionario.CrearABB[int, int](cmpInt)

	claves := []int{5, 3, 8, 1, 6, 4, 7, 2, 9, 10}

	for _, clave := range claves {
		abb.Guardar(clave, clave)
	}

	desde := 3
	hasta := 7
	suma := 0
	abb.IterarRango(&desde, &hasta, func(_ int, dato int) bool {
		suma += dato
		return suma < 10 // Cortar cuando la suma sea mayor o igual a 10
	})

	require.EqualValues(t, 12, suma) // 3 + 4 + 5 = 12
}

func TestIterarInternoCondicionCorteRecursivo(t *testing.T) {
	t.Log("Prueba el iterador interno con condición de corte buscando que si !visitar, esto se extienda recursivamente")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)

	dic.Guardar("b", 2)
	dic.Guardar("a", 1)
	dic.Guardar("c", 3)

	var clavesConcatenadas string
	visitar := func(clave string, valor int) bool {
		clavesConcatenadas += clave
		return clave != "a"
	}

	dic.Iterar(visitar)
	require.EqualValues(t, "a", clavesConcatenadas)
}

func TestIteradorInternoRangoVacio(t *testing.T) {
	t.Log("Valida que el iterador interno no itere cuando el rango es vacío")
	abb := TDADiccionario.CrearABB[int, int](cmpInt)

	claves := []int{5, 3, 8, 1, 6, 4, 7, 2, 9, 10}

	for _, clave := range claves {
		abb.Guardar(clave, clave)
	}

	desde := 11
	hasta := 15
	suma := 0
	abb.IterarRango(&desde, &hasta, func(_ int, dato int) bool {
		suma += dato
		return true
	})

	require.EqualValues(t, 0, suma)
}

func TestIteradorInternoRangoUnElemento(t *testing.T) {
	t.Log("Valida que el iterador interno funcione con un rango que contiene un solo elemento")
	abb := TDADiccionario.CrearABB[int, int](cmpInt)

	claves := []int{5, 3, 8, 1, 6, 4, 7, 2, 9, 10}

	for _, clave := range claves {
		abb.Guardar(clave, clave)
	}

	desde := 5
	hasta := 5
	suma := 0
	abb.IterarRango(&desde, &hasta, func(_ int, dato int) bool {
		suma += dato
		return true
	})

	require.EqualValues(t, 5, suma)
}

func TestIteradorInternoRangoInverso(t *testing.T) {
	t.Log("Valida que el iterador interno no itere cuando el rango es inverso")
	abb := TDADiccionario.CrearABB[int, int](cmpInt)

	claves := []int{5, 3, 8, 1, 6, 4, 7, 2, 9, 10}

	for _, clave := range claves {
		abb.Guardar(clave, clave)
	}

	desde := 7
	hasta := 3
	suma := 0
	abb.IterarRango(&desde, &hasta, func(_ int, dato int) bool {
		suma += dato
		return true
	})

	require.EqualValues(t, 0, suma)
}

/* Tests iterador externo */

func TestIterarDicOrdenadoVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDicOrdenadoIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente siguen un recorrido in-order")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, primer_valor := iter.VerActual()
	require.EqualValues(t, clave1, primero)
	require.EqualValues(t, valor1, primer_valor)

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.EqualValues(t, clave2, segundo)
	require.EqualValues(t, valor2, segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, tercer_valor := iter.VerActual()
	require.EqualValues(t, clave3, tercero)
	require.EqualValues(t, valor3, tercer_valor)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorNoLlegaAFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.EqualValues(t, primero, claves[0])
	require.EqualValues(t, segundo, claves[1])
	require.EqualValues(t, tercero, claves[2])
}

func correrPruebasVolumenIterador(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)

	claves := make([]string, n)
	valores := make([]int, n)

	/* Guarda "n" parejas en los arreglos */
	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = fmt.Sprintf("%08d", i)
	}

	mezclarClaves(claves) // Mezcla las claves para insertarlas en un orden aleatorio

	/* Inserta "n" parejas en el diccionario */
	for i := 0; i < n; i++ {
		dic.Guardar(claves[i], &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterador()
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	var clave string
	var valor *int
	var anterior *int

	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n
		if anterior != nil && *anterior != n {
			require.True(b, cmpInt(*anterior, *valor) < 0, "No se recorrió en orden")
		}
		anterior = valor
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}

func BenchmarkIteradorOrdenado(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMANIOS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				correrPruebasVolumenIterador(b, n)
			}
		})
	}
}

func TestIteradorRango(t *testing.T) {
	t.Log("Prueba el funcionamiento de el iterador externo con un rango definido.")
	dic := TDADiccionario.CrearABB[int, string](cmpInt)
	claves := []int{100, 50, 25, 10, 30, 75, 60, 80, 150, 125, 110, 130, 200, 190, 300}
	for _, clave := range claves {
		dic.Guardar(clave, "")
	}
	desde := 60
	hasta := 120
	iter := dic.IteradorRango(&desde, &hasta)

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.EqualValues(t, 60, primero)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	segundo, _ := iter.VerActual()
	require.EqualValues(t, 75, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	tercero, _ := iter.VerActual()
	require.EqualValues(t, 80, tercero)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	cuarto, _ := iter.VerActual()
	require.EqualValues(t, 100, cuarto)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	quinto, _ := iter.VerActual()
	require.EqualValues(t, 110, quinto)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorRangoSinDesde(t *testing.T) {
	t.Log("Prueba que el iterador externo con desde igual a nil itere desde la primera clave.")
	dic := TDADiccionario.CrearABB[int, string](cmpInt)
	claves := []int{100, 50, 25, 10, 30, 75, 60, 80, 150, 125, 110, 130, 200, 190, 300}
	for _, clave := range claves {
		dic.Guardar(clave, "")
	}
	hasta := 120
	iter := dic.IteradorRango(nil, &hasta)

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.EqualValues(t, 10, primero)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	segundo, _ := iter.VerActual()
	require.EqualValues(t, 25, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	tercero, _ := iter.VerActual()
	require.EqualValues(t, 30, tercero)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	cuarto, _ := iter.VerActual()
	require.EqualValues(t, 50, cuarto)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	quinto, _ := iter.VerActual()
	require.EqualValues(t, 60, quinto)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	sexto, _ := iter.VerActual()
	require.EqualValues(t, 75, sexto)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	septimo, _ := iter.VerActual()
	require.EqualValues(t, 80, septimo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	octavo, _ := iter.VerActual()
	require.EqualValues(t, 100, octavo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	noveno, _ := iter.VerActual()
	require.EqualValues(t, 110, noveno)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorRangoSinHasta(t *testing.T) {
	t.Log("Prueba que el iterador externo con hasta igual a nil itere hasta la última clave.")
	dic := TDADiccionario.CrearABB[int, string](cmpInt)
	claves := []int{100, 50, 25, 10, 30, 75, 60, 80, 150, 125, 110, 130, 200, 190, 300}
	for _, clave := range claves {
		dic.Guardar(clave, "")
	}
	desde := 120
	iter := dic.IteradorRango(&desde, nil)

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.EqualValues(t, 125, primero)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	segundo, _ := iter.VerActual()
	require.EqualValues(t, 130, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	tercero, _ := iter.VerActual()
	require.EqualValues(t, 150, tercero)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	cuarto, _ := iter.VerActual()
	require.EqualValues(t, 190, cuarto)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	quinto, _ := iter.VerActual()
	require.EqualValues(t, 200, quinto)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	sexto, _ := iter.VerActual()
	require.EqualValues(t, 300, sexto)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorRangoSinRango(t *testing.T) {
	t.Log("Prueba que el iterador externo con desde y hasta iguales a nil se comporta como el iterador sin rango.")
	dic := TDADiccionario.CrearABB[int, string](cmpInt)
	claves := []int{100, 50, 25, 75, 150, 125}
	for _, clave := range claves {
		dic.Guardar(clave, "")
	}
	iter := dic.IteradorRango(nil, nil)

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.EqualValues(t, 25, primero)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	segundo, _ := iter.VerActual()
	require.EqualValues(t, 50, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	tercero, _ := iter.VerActual()
	require.EqualValues(t, 75, tercero)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	cuarto, _ := iter.VerActual()
	require.EqualValues(t, 100, cuarto)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	quinto, _ := iter.VerActual()
	require.EqualValues(t, 125, quinto)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	sexto, _ := iter.VerActual()
	require.EqualValues(t, 150, sexto)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorExternoVacio(t *testing.T) {
	t.Log("Prueba el iterador externo con rango con un diccionario vacío")
	abb := TDADiccionario.CrearABB[int, int](cmpInt)
	desde := 5
	hasta := 6
	iter := abb.IteradorRango(&desde, &hasta)
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorInternoOrdenados(t *testing.T) {
	t.Log("Valida que el iterador interno itere en orden aunque los elementos se hayan insertado desordenados")
	abb := TDADiccionario.CrearABB[int, int](cmpInt)

	elems := make([]int, 400)

	for i := 0; i < 400; i++ {
		elems[i] = i
	}

	elems = mezclarClaves(elems)

	for _, elem := range elems {
		abb.Guardar(elem, elem)
	}

	orden := []int{}
	abb.Iterar(func(clave int, dato int) bool {
		orden = append(orden, clave)
		return true
	})

	slices.Sort(elems)

	require.EqualValues(t, elems, orden)
}

func TestIteradorRangoOrdenados(t *testing.T) {
	t.Log("Valida que el iterador interno por rangos itere en orden aunque los elementos se hayan insertado desordenados")
	abb := TDADiccionario.CrearABB[int, int](cmpInt)

	elems := make([]int, 400)

	for i := 0; i < 400; i++ {
		elems[i] = i
	}

	elems = mezclarClaves(elems)

	for _, elem := range elems {
		abb.Guardar(elem, elem)
	}

	orden := []int{}
	ini := 100
	fin := 300
	abb.IterarRango(&ini, &fin, func(clave int, dato int) bool {
		orden = append(orden, clave)
		return true
	})

	slices.Sort(elems)

	require.EqualValues(t, elems[ini:fin+1], orden)
}

func TestIteradorExternoOrdenados(t *testing.T) {
	t.Log("Valida que el iterador externo itere en orden aunque los elementos se hayan insertado desordenados")
	abb := TDADiccionario.CrearABB[int, int](cmpInt)

	elems := make([]int, 400)

	for i := 0; i < 400; i++ {
		elems[i] = i
	}

	elems = mezclarClaves(elems)

	for _, elem := range elems {
		abb.Guardar(elem, elem)
	}

	orden := []int{}
	iter := abb.Iterador()
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		orden = append(orden, clave)
		iter.Siguiente()
	}

	slices.Sort(elems)

	require.EqualValues(t, elems, orden)
}

func TestIteradorExternoRangosOrdenados(t *testing.T) {
	t.Log("Valida que el iterador externo itere en orden aunque los elementos se hayan insertado desordenados")
	abb := TDADiccionario.CrearABB[int, int](cmpInt)

	elems := make([]int, 400)

	for i := 0; i < 400; i++ {
		elems[i] = i
	}

	elems = mezclarClaves(elems)

	for _, elem := range elems {
		abb.Guardar(elem, elem)
	}

	orden := []int{}
	ini := 40
	fin := 250
	iter := abb.IteradorRango(&ini, &fin)
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		orden = append(orden, clave)
		iter.Siguiente()
	}

	slices.Sort(elems)

	require.EqualValues(t, elems[ini:fin+1], orden)
}
