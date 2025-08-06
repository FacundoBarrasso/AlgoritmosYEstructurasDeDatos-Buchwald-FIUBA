package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

// Las pruebas deberán verificar que:

// 1. Se pueda crear una Pila vacía, y esta se comporta como tal.
// 2. Se puedan apilar elementos, que al desapilarlos se mantenga el invariante de pila (que esta es LIFO). Probar con elementos diferentes, y ver que salgan en el orden deseado.
// 3. Prueba de volumen: Se pueden apilar muchos elementos (1000, 10000 elementos, o el volumen que corresponda): hacer crecer la pila, y desapilar elementos hasta que esté vacía, comprobando que siempre cumpla el invariante. Recordar no apilar siempre lo mismo, validar que se cumpla siempre que el tope de la pila sea el correcto paso a paso, y que el nuevo tope después de cada desapilar también sea el correcto.
// 4. Condición de borde 1: comprobar que al desapilar hasta que está vacía hace que la pila se comporte como recién creada.
// 5. Condición de borde 2: las acciones de desapilar y ver_tope en una pila recién creada son inválidas.
// 6. Condición de borde 3: la acción de esta_vacía en una pila recién creada es verdadero.
// 7. Condición de borde 4: las acciones de desapilar y ver_tope en una pila a la que se le apiló y desapiló hasta estar vacía son inválidas.
// 8. Probar apilar diferentes tipos de datos: probar con una pila de enteros, con una pila de cadenas, etc…

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestApilarYDesapilarLIFO(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	require.False(t, pila.EstaVacia())
	require.Equal(t, 3, pila.VerTope())
	require.Equal(t, 3, pila.Desapilar())
	require.Equal(t, 2, pila.Desapilar())
	require.Equal(t, 1, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestPruebaDeVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 100000; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
	}
	for i := 99999; i >= 0; i-- {
		require.Equal(t, i, pila.Desapilar())
	}
	require.True(t, pila.EstaVacia())
}

func TestCondicionDeBorde1(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Desapilar()
	pila.Apilar(2)
	pila.Desapilar()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	// Quise comparar con require.Equal que las salidas de la pila normal y una pila recien creada sean iguales, para no tener
	//  practicamente las mismas pruebas que en el TestCondicionDeBorde4, pero no logré que funcione correctamente.
}

func TestCondicionDeBorde2(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestCondicionDeBorde3(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
}

func TestCondicionDeBorde4(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Desapilar()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestApilarDiferentesTipos(t *testing.T) {
	pilaInt := TDAPila.CrearPilaDinamica[int]()
	pilaInt.Apilar(1)
	require.Equal(t, 1, pilaInt.VerTope())
	require.Equal(t, 1, pilaInt.Desapilar())

	pilaString := TDAPila.CrearPilaDinamica[string]()
	pilaString.Apilar("hola")
	require.Equal(t, "hola", pilaString.VerTope())
	require.Equal(t, "hola", pilaString.Desapilar())

	pilaFloat := TDAPila.CrearPilaDinamica[float64]()
	pilaFloat.Apilar(1.5)
	require.Equal(t, 1.5, pilaFloat.VerTope())
	require.Equal(t, 1.5, pilaFloat.Desapilar())

	pilaBool := TDAPila.CrearPilaDinamica[bool]()
	pilaBool.Apilar(true)
	require.Equal(t, true, pilaBool.VerTope())
	require.Equal(t, true, pilaBool.Desapilar())

	pilaRune := TDAPila.CrearPilaDinamica[rune]()
	pilaRune.Apilar('a')
	require.Equal(t, 'a', pilaRune.VerTope())
	require.Equal(t, 'a', pilaRune.Desapilar())

	pilaComplex := TDAPila.CrearPilaDinamica[complex128]()
	pilaComplex.Apilar(1 + 2i)
	require.Equal(t, 1+2i, pilaComplex.VerTope())
	require.Equal(t, 1+2i, pilaComplex.Desapilar())
}
