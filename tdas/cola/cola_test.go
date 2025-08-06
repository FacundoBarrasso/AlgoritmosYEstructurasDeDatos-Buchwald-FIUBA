package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

// Las pruebas deberán verificar que:

// 1. Se pueda crear una Cola vacía, y esta se comporta como tal.
// 2. Se puedan encolar elementos, que al desencolarlos se mantenga el invariante de cola (que esta es FIFO). Probar con elementos diferentes, y ver que salgan en el orden deseado.
// 3. Prueba de volumen: Se pueden encolar muchos elementos (1000, 10000 elementos, o el volumen que corresponda): hacer crecer la cola, y desencolar elementos hasta que esté vacía, comprobando que siempre cumpla el invariante. Recordar no encolar siempre lo mismo, validar que se cumpla siempre que el primero de la cola sea el correcto paso a paso, y que el nuevo primero después de cada desencolar también sea el correcto.
// 4. Condición de borde 1: comprobar que al desencolar hasta que está vacía hace que la cola se comporte como recién creada.
// 5. Condición de borde 2: las acciones de desencolar y ver_primero en una cola recién creada son inválidas.
// 6. Condición de borde 3: la acción de esta_vacía en una cola recién creada es verdadero.
// 7. Condición de borde 4: las acciones de desencolar y ver_primero en una cola a la que se le encoló y desencoló hasta estar vacía son inválidas.
// 8. Probar encolar diferentes tipos de datos: probar con una cola de enteros, con una cola de cadenas, etc…

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestEncolarYDesencolarFIFO(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	require.Equal(t, 1, cola.VerPrimero())
	cola.Encolar(2)
	require.Equal(t, 1, cola.VerPrimero())
	cola.Encolar(3)
	require.Equal(t, 1, cola.VerPrimero())
	require.False(t, cola.EstaVacia())
	require.Equal(t, 1, cola.VerPrimero())
	require.Equal(t, 1, cola.Desencolar())
	require.Equal(t, 2, cola.VerPrimero())
	require.Equal(t, 2, cola.Desencolar())
	require.Equal(t, 3, cola.VerPrimero())
	require.Equal(t, 3, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestPruebaDeVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 100000; i++ {
		cola.Encolar(i)
		require.Equal(t, 0, cola.VerPrimero()) // El primero no deberia cambiar al encolar
	}
	j := 0
	for i := 99999; i >= 0; i-- {
		require.Equal(t, j, cola.VerPrimero()) // El primero deberia cambiar al desencolar
		require.Equal(t, j, cola.Desencolar())
		j++
	}
	require.True(t, cola.EstaVacia())
}

func TestCondicionDeBorde1(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Desencolar()
	cola.Desencolar()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestCondicionDeBorde2(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestCondicionDeBorde3(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
}

func TestCondicionDeBorde4(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Desencolar()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestEncolarDiferentesTipos(t *testing.T) {
	colaInt := TDACola.CrearColaEnlazada[int]()
	colaInt.Encolar(1)
	require.Equal(t, 1, colaInt.VerPrimero())
	require.Equal(t, 1, colaInt.Desencolar())

	colaString := TDACola.CrearColaEnlazada[string]()
	colaString.Encolar("hola")
	require.Equal(t, "hola", colaString.VerPrimero())
	require.Equal(t, "hola", colaString.Desencolar())

	colaFloat := TDACola.CrearColaEnlazada[float64]()
	colaFloat.Encolar(1.5)
	require.Equal(t, 1.5, colaFloat.VerPrimero())
	require.Equal(t, 1.5, colaFloat.Desencolar())

	colaBool := TDACola.CrearColaEnlazada[bool]()
	colaBool.Encolar(true)
	require.Equal(t, true, colaBool.VerPrimero())
	require.Equal(t, true, colaBool.Desencolar())

	colaRune := TDACola.CrearColaEnlazada[rune]()
	colaRune.Encolar('a')
	require.Equal(t, 'a', colaRune.VerPrimero())
	require.Equal(t, 'a', colaRune.Desencolar())

	colaComplex := TDACola.CrearColaEnlazada[complex128]()
	colaComplex.Encolar(1 + 2i)
	require.Equal(t, 1+2i, colaComplex.VerPrimero())
	require.Equal(t, 1+2i, colaComplex.Desencolar())
}
