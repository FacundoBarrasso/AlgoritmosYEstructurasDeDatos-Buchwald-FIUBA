package cola_prioridad_test

import (
	"fmt"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"math/rand"

	"github.com/stretchr/testify/require"
)

// Funciones de comparación para diferentes tipos de datos
func compararMaximos(a, b int) int {
	return a - b
}

func compararMinimos(a, b int) int {
	return b - a
}

func compararStrings(a, b string) int {
	return strings.Compare(a, b)
}

type Persona struct {
	nombre string
	edad   int
}

func compararPersonasEdadMax(a, b Persona) int {
	return compararMaximos(a.edad, b.edad)
}

func compararPersonasEdadMin(a, b Persona) int {
	return compararMinimos(a.edad, b.edad)
}

func compararPersonasNombreMax(a, b Persona) int {
	return compararStrings(a.nombre, b.nombre)
}

func compararPersonasNombreMin(a, b Persona) int {
	return compararStrings(b.nombre, a.nombre)
}

var TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

func testHeapVacio[T any](t *testing.T, heap TDAHeap.ColaPrioridad[T]) {
	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapMaxVacio(t *testing.T) {
	t.Log("Prueba heap de máximos vacío")
	heap := TDAHeap.CrearHeap(compararMaximos)
	testHeapVacio(t, heap)
}

func TestHeapMinVacio(t *testing.T) {
	t.Log("Prueba heap de mínimos vacío")
	heap := TDAHeap.CrearHeap(compararMinimos)
	testHeapVacio(t, heap)
}

func TestHeapMaxEncolarParDesencolar(t *testing.T) {
	t.Log("Prueba heap de máximos con cantidad par de elementos")
	heap := TDAHeap.CrearHeap(compararMaximos)

	elementos := []int{5, 2, 7, 1, 3, 6}
	for _, elem := range elementos {
		heap.Encolar(elem)
	}

	require.Equal(t, 6, heap.Cantidad())
	require.Equal(t, 7, heap.VerMax())

	require.Equal(t, 7, heap.Desencolar())
	require.Equal(t, 5, heap.Cantidad())
	require.Equal(t, 6, heap.VerMax())

	require.Equal(t, 6, heap.Desencolar())
	require.Equal(t, 4, heap.Cantidad())
	require.Equal(t, 5, heap.VerMax())

	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, 3, heap.VerMax())

	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 2, heap.Cantidad())
	require.Equal(t, 2, heap.VerMax())

	require.Equal(t, 2, heap.Desencolar())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 1, heap.VerMax())

	require.Equal(t, 1, heap.Desencolar())
	testHeapVacio(t, heap)
}

func TestHeapMinEncolarParDesencolar(t *testing.T) {
	t.Log("Prueba heap de mínimos con cantidad par de elementos")
	heap := TDAHeap.CrearHeap(compararMinimos)

	elementos := []int{5, 2, 7, 1, 3, 6}
	for _, elem := range elementos {
		heap.Encolar(elem)
	}

	require.Equal(t, 6, heap.Cantidad())
	require.Equal(t, 1, heap.VerMax())

	require.Equal(t, 1, heap.Desencolar())
	require.Equal(t, 5, heap.Cantidad())
	require.Equal(t, 2, heap.VerMax())

	require.Equal(t, 2, heap.Desencolar())
	require.Equal(t, 4, heap.Cantidad())
	require.Equal(t, 3, heap.VerMax())

	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, 5, heap.VerMax())

	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 2, heap.Cantidad())
	require.Equal(t, 6, heap.VerMax())

	require.Equal(t, 6, heap.Desencolar())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 7, heap.VerMax())

	require.Equal(t, 7, heap.Desencolar())
	testHeapVacio(t, heap)
}

func TestHeapMaxEncolarImparDesencolar(t *testing.T) {
	t.Log("Prueba heap de máximos con cantidad impar de elementos")
	heap := TDAHeap.CrearHeap(compararMaximos)

	elementos := []int{5, 2, 7, 1, 3}
	for _, elem := range elementos {
		heap.Encolar(elem)
	}

	require.Equal(t, 5, heap.Cantidad())
	require.Equal(t, 7, heap.VerMax())

	require.Equal(t, 7, heap.Desencolar())
	require.Equal(t, 4, heap.Cantidad())
	require.Equal(t, 5, heap.VerMax())

	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, 3, heap.VerMax())

	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 2, heap.Cantidad())
	require.Equal(t, 2, heap.VerMax())

	require.Equal(t, 2, heap.Desencolar())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 1, heap.VerMax())

	require.Equal(t, 1, heap.Desencolar())
	testHeapVacio(t, heap)
}

func TestHeapMinEncolarImparDesencolar(t *testing.T) {
	t.Log("Prueba heap de mínimos con cantidad impar de elementos")
	heap := TDAHeap.CrearHeap(compararMinimos)

	elementos := []int{5, 2, 7, 1, 3}
	for _, elem := range elementos {
		heap.Encolar(elem)
	}

	require.Equal(t, 5, heap.Cantidad())
	require.Equal(t, 1, heap.VerMax())

	require.Equal(t, 1, heap.Desencolar())
	require.Equal(t, 4, heap.Cantidad())
	require.Equal(t, 2, heap.VerMax())

	require.Equal(t, 2, heap.Desencolar())
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, 3, heap.VerMax())

	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 2, heap.Cantidad())
	require.Equal(t, 5, heap.VerMax())

	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 7, heap.VerMax())

	require.Equal(t, 7, heap.Desencolar())
	testHeapVacio(t, heap)
}

func TestCrearHeapArrMax(t *testing.T) {
	t.Log("Prueba crear heap de máximos a partir de un arreglo")
	arreglo := []int{5, 2, 7, 1, 3}
	heap := TDAHeap.CrearHeapArr(arreglo, compararMaximos)

	require.Equal(t, 5, heap.Cantidad())
	require.Equal(t, 7, heap.VerMax())
	require.Equal(t, 7, heap.Desencolar())
	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 2, heap.Desencolar())
	require.Equal(t, 1, heap.Desencolar())
	testHeapVacio(t, heap)
}

func TestCrearHeapArrMin(t *testing.T) {
	t.Log("Prueba crear heap de mínimos a partir de un arreglo")
	arreglo := []int{5, 2, 7, 1, 3}
	heap := TDAHeap.CrearHeapArr(arreglo, compararMinimos)

	require.Equal(t, 5, heap.Cantidad())
	require.Equal(t, 1, heap.VerMax())
	require.Equal(t, 1, heap.Desencolar())
	require.Equal(t, 2, heap.Desencolar())
	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 7, heap.Desencolar())
	testHeapVacio(t, heap)
}
func TestHeapStrings(t *testing.T) {
	t.Log("Prueba heap con strings")
	heap := TDAHeap.CrearHeap(compararStrings)

	palabras := []string{"casa", "perro", "avion", "zebra", "barco"}
	for _, palabra := range palabras {
		heap.Encolar(palabra)
	}

	require.Equal(t, 5, heap.Cantidad())
	require.Equal(t, "zebra", heap.VerMax())
	require.Equal(t, "zebra", heap.Desencolar())
	require.Equal(t, "perro", heap.Desencolar())
	require.Equal(t, "casa", heap.Desencolar())
	require.Equal(t, "barco", heap.Desencolar())
	require.Equal(t, "avion", heap.Desencolar())
	testHeapVacio(t, heap)
}

func TestHeapPersonasEdadMax(t *testing.T) {
	t.Log("Prueba heap de máximos con estructuras Persona ordenadas por edad (de el más grande al más chico)")
	heap := TDAHeap.CrearHeap(compararPersonasEdadMax)

	personas := []Persona{
		{"Maria", 30},
		{"Pedro", 20},
		{"Ana", 35},
		{"Luis", 28},
	}

	for _, persona := range personas {
		heap.Encolar(persona)
	}

	require.Equal(t, 4, heap.Cantidad())
	maxPersona := heap.VerMax()
	require.Equal(t, "Ana", maxPersona.nombre)
	require.Equal(t, 35, maxPersona.edad)

	persona := heap.Desencolar()
	require.Equal(t, 35, persona.edad)
	persona = heap.Desencolar()
	require.Equal(t, 30, persona.edad)
	persona = heap.Desencolar()
	require.Equal(t, 28, persona.edad)
	persona = heap.Desencolar()
	require.Equal(t, 20, persona.edad)
	testHeapVacio(t, heap)
}

func TestHeapPersonasEdadMin(t *testing.T) {
	t.Log("Prueba heap de mínimos con estructuras Persona ordenadas por edad (de el más chico al más grande)")
	heap := TDAHeap.CrearHeap(compararPersonasEdadMin)

	personas := []Persona{
		{"Maria", 30},
		{"Pedro", 20},
		{"Ana", 35},
		{"Luis", 28},
	}

	for _, persona := range personas {
		heap.Encolar(persona)
	}

	require.Equal(t, 4, heap.Cantidad())
	maxPersona := heap.VerMax()
	require.Equal(t, "Pedro", maxPersona.nombre)
	require.Equal(t, 20, maxPersona.edad)

	persona := heap.Desencolar()
	require.Equal(t, 20, persona.edad)
	persona = heap.Desencolar()
	require.Equal(t, 28, persona.edad)
	persona = heap.Desencolar()
	require.Equal(t, 30, persona.edad)
	persona = heap.Desencolar()
	require.Equal(t, 35, persona.edad)
	testHeapVacio(t, heap)
}

func TestHeapPersonasNombreMin(t *testing.T) {
	t.Log("Prueba heap de mínimos con estructuras Persona ordenadas por nombre (A-Z)")
	heap := TDAHeap.CrearHeap(compararPersonasNombreMin)

	personas := []Persona{
		{"Juan", 25},
		{"Maria", 30},
		{"Pedro", 20},
		{"Ana", 35},
		{"Luis", 28},
	}

	for _, persona := range personas {
		heap.Encolar(persona)
	}

	require.Equal(t, 5, heap.Cantidad())
	maxPersona := heap.VerMax()
	require.Equal(t, "Ana", maxPersona.nombre)

	persona := heap.Desencolar()
	require.Equal(t, "Ana", persona.nombre)
	persona = heap.Desencolar()
	require.Equal(t, "Juan", persona.nombre)
	persona = heap.Desencolar()
	require.Equal(t, "Luis", persona.nombre)
	persona = heap.Desencolar()
	require.Equal(t, "Maria", persona.nombre)
	persona = heap.Desencolar()
	require.Equal(t, "Pedro", persona.nombre)
	testHeapVacio(t, heap)
}

func TestHeapPersonasNombreMax(t *testing.T) {
	t.Log("Prueba heap de máximos con estructuras Persona ordenadas por nombre (Z-A)")
	heap := TDAHeap.CrearHeap(compararPersonasNombreMax)
	personas := []Persona{
		{"Juan", 25},
		{"Maria", 30},
		{"Pedro", 20},
		{"Ana", 35},
		{"Luis", 28},
	}

	for _, persona := range personas {
		heap.Encolar(persona)
	}

	require.Equal(t, 5, heap.Cantidad())
	maxPersona := heap.VerMax()
	require.Equal(t, "Pedro", maxPersona.nombre)

	persona := heap.Desencolar()
	require.Equal(t, "Pedro", persona.nombre)
	persona = heap.Desencolar()
	require.Equal(t, "Maria", persona.nombre)
	persona = heap.Desencolar()
	require.Equal(t, "Luis", persona.nombre)
	persona = heap.Desencolar()
	require.Equal(t, "Juan", persona.nombre)
	persona = heap.Desencolar()
	require.Equal(t, "Ana", persona.nombre)
	testHeapVacio(t, heap)
}

func TestHeapMaxConDuplicados(t *testing.T) {
	t.Log("Prueba heap de máximos con elementos duplicados")
	heap := TDAHeap.CrearHeap(compararMaximos)

	elementos := []int{5, 2, 7, 1, 3, 7, 2}
	for _, elem := range elementos {
		heap.Encolar(elem)
	}

	require.Equal(t, 7, heap.Cantidad())
	require.Equal(t, 7, heap.VerMax())

	require.Equal(t, 7, heap.Desencolar())
	require.Equal(t, 6, heap.Cantidad())
	require.Equal(t, 7, heap.VerMax())

	require.Equal(t, 7, heap.Desencolar())
	require.Equal(t, 5, heap.Cantidad())
	require.Equal(t, 5, heap.VerMax())

	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 4, heap.Cantidad())
	require.Equal(t, 3, heap.VerMax())

	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, 2, heap.VerMax())

	require.Equal(t, 2, heap.Desencolar())
	require.Equal(t, 2, heap.Cantidad())
	require.Equal(t, 2, heap.VerMax())

	require.Equal(t, 2, heap.Desencolar())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 1, heap.VerMax())

	require.Equal(t, 1, heap.Desencolar())
	testHeapVacio(t, heap)
}

func TestHeapMinConDuplicados(t *testing.T) {
	t.Log("Prueba heap de mínimos con elementos duplicados")
	heap := TDAHeap.CrearHeap(compararMinimos)

	elementos := []int{5, 2, 7, 1, 3, 7, 2}
	for _, elem := range elementos {
		heap.Encolar(elem)
	}

	require.Equal(t, 7, heap.Cantidad())
	require.Equal(t, 1, heap.VerMax())

	require.Equal(t, 1, heap.Desencolar())
	require.Equal(t, 6, heap.Cantidad())
	require.Equal(t, 2, heap.VerMax())

	require.Equal(t, 2, heap.Desencolar())
	require.Equal(t, 5, heap.Cantidad())
	require.Equal(t, 2, heap.VerMax())

	require.Equal(t, 2, heap.Desencolar())
	require.Equal(t, 4, heap.Cantidad())
	require.Equal(t, 3, heap.VerMax())

	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, 5, heap.VerMax())

	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 2, heap.Cantidad())
	require.Equal(t, 7, heap.VerMax())

	require.Equal(t, 7, heap.Desencolar())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 7, heap.VerMax())

	require.Equal(t, 7, heap.Desencolar())
	testHeapVacio(t, heap)
}

func TestHeapStringsConDuplicados(t *testing.T) {
	t.Log("Prueba heap con strings duplicados")
	heap := TDAHeap.CrearHeap(compararStrings)

	palabras := []string{"casa", "perro", "avion", "zebra", "barco", "perro"}
	for _, palabra := range palabras {
		heap.Encolar(palabra)
	}

	require.Equal(t, 6, heap.Cantidad())
	require.Equal(t, "zebra", heap.VerMax())
	require.Equal(t, "zebra", heap.Desencolar())
	require.Equal(t, "perro", heap.VerMax())
	require.Equal(t, "perro", heap.Desencolar())
	require.Equal(t, "perro", heap.VerMax())
	require.Equal(t, "perro", heap.Desencolar())
	require.Equal(t, "casa", heap.VerMax())
	require.Equal(t, "casa", heap.Desencolar())
	require.Equal(t, "barco", heap.VerMax())
	require.Equal(t, "barco", heap.Desencolar())
	require.Equal(t, "avion", heap.VerMax())
	require.Equal(t, "avion", heap.Desencolar())
	testHeapVacio(t, heap)
}

func TestHeapMaxConUnElemento(t *testing.T) {
	t.Log("Prueba heap de máximos con un solo elemento")
	heap := TDAHeap.CrearHeap(compararMaximos)

	heap.Encolar(10)
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 10, heap.VerMax())
	require.Equal(t, 10, heap.Desencolar())
	testHeapVacio(t, heap)
}

func TestHeapMinConUnElemento(t *testing.T) {
	t.Log("Prueba heap de mínimos con un solo elemento")
	heap := TDAHeap.CrearHeap(compararMinimos)

	heap.Encolar(10)
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 10, heap.VerMax())
	require.Equal(t, 10, heap.Desencolar())
	testHeapVacio(t, heap)
}

func TestHeapMaxConElementosIguales(t *testing.T) {
	t.Log("Prueba heap de máximos con elementos iguales")
	heap := TDAHeap.CrearHeap(compararMaximos)

	elementos := []int{5, 5, 5, 5, 5}
	for _, elem := range elementos {
		heap.Encolar(elem)
	}

	require.Equal(t, 5, heap.Cantidad())
	require.Equal(t, 5, heap.VerMax())

	for i := 0; i < 5; i++ {
		require.Equal(t, 5, heap.Desencolar())
	}
	testHeapVacio(t, heap)
}

func TestHeapMinConElementosIguales(t *testing.T) {
	t.Log("Prueba heap de mínimos con elementos iguales")
	heap := TDAHeap.CrearHeap(compararMinimos)

	elementos := []int{5, 5, 5, 5, 5}
	for _, elem := range elementos {
		heap.Encolar(elem)
	}

	require.Equal(t, 5, heap.Cantidad())
	require.Equal(t, 5, heap.VerMax())

	for i := 0; i < 5; i++ {
		require.Equal(t, 5, heap.Desencolar())
	}
	testHeapVacio(t, heap)
}

func TestHeapMaxConElementosNegativos(t *testing.T) {
	t.Log("Prueba heap de máximos con elementos negativos")
	heap := TDAHeap.CrearHeap(compararMaximos)

	elementos := []int{-5, -2, -7, -1, -3}
	for _, elem := range elementos {
		heap.Encolar(elem)
	}

	require.Equal(t, 5, heap.Cantidad())
	require.Equal(t, -1, heap.VerMax())

	require.Equal(t, -1, heap.Desencolar())
	require.Equal(t, -2, heap.VerMax())
	require.Equal(t, -2, heap.Desencolar())
	require.Equal(t, -3, heap.VerMax())
	require.Equal(t, -3, heap.Desencolar())
	require.Equal(t, -5, heap.VerMax())
	require.Equal(t, -5, heap.Desencolar())
	require.Equal(t, -7, heap.VerMax())
	require.Equal(t, -7, heap.Desencolar())
	testHeapVacio(t, heap)
}

func TestHeapMinConElementosNegativos(t *testing.T) {
	t.Log("Prueba heap de mínimos con elementos negativos")
	heap := TDAHeap.CrearHeap(compararMinimos)

	elementos := []int{-5, -2, -7, -1, -3}
	for _, elem := range elementos {
		heap.Encolar(elem)
	}

	require.Equal(t, 5, heap.Cantidad())
	require.Equal(t, -7, heap.VerMax())

	require.Equal(t, -7, heap.Desencolar())
	require.Equal(t, -5, heap.VerMax())
	require.Equal(t, -5, heap.Desencolar())
	require.Equal(t, -3, heap.VerMax())
	require.Equal(t, -3, heap.Desencolar())
	require.Equal(t, -2, heap.VerMax())
	require.Equal(t, -2, heap.Desencolar())
	require.Equal(t, -1, heap.VerMax())
	require.Equal(t, -1, heap.Desencolar())
	testHeapVacio(t, heap)
}

func TestHeapMaxConElementosMixtos(t *testing.T) {
	t.Log("Prueba heap de máximos con elementos positivos y negativos")
	heap := TDAHeap.CrearHeap(compararMaximos)

	elementos := []int{5, -2, 7, -1, 3, 0}
	for _, elem := range elementos {
		heap.Encolar(elem)
	}

	require.Equal(t, 6, heap.Cantidad())
	require.Equal(t, 7, heap.VerMax())

	require.Equal(t, 7, heap.Desencolar())
	require.Equal(t, 5, heap.Cantidad())
	require.Equal(t, 5, heap.VerMax())

	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 4, heap.Cantidad())
	require.Equal(t, 3, heap.VerMax())

	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, 0, heap.VerMax())

	require.Equal(t, 0, heap.Desencolar())
	require.Equal(t, 2, heap.Cantidad())
	require.Equal(t, -1, heap.VerMax())

	require.Equal(t, -1, heap.Desencolar())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, -2, heap.VerMax())

	require.Equal(t, -2, heap.Desencolar())
	testHeapVacio(t, heap)
}

func TestHeapMinConElementosMixtos(t *testing.T) {
	t.Log("Prueba heap de mínimos con elementos elementos positivos y negativos")
	heap := TDAHeap.CrearHeap(compararMinimos)

	elementos := []int{5, -2, 7, -1, 3, 0}
	for _, elem := range elementos {
		heap.Encolar(elem)
	}

	require.Equal(t, 6, heap.Cantidad())
	require.Equal(t, -2, heap.VerMax())

	require.Equal(t, -2, heap.Desencolar())
	require.Equal(t, 5, heap.Cantidad())
	require.Equal(t, -1, heap.VerMax())

	require.Equal(t, -1, heap.Desencolar())
	require.Equal(t, 4, heap.Cantidad())
	require.Equal(t, 0, heap.VerMax())

	require.Equal(t, 0, heap.Desencolar())
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, 3, heap.VerMax())

	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 2, heap.Cantidad())
	require.Equal(t, 5, heap.VerMax())

	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 7, heap.VerMax())

	require.Equal(t, 7, heap.Desencolar())
	testHeapVacio(t, heap)
}

func TestHeapSortMax(t *testing.T) {
	t.Log("Prueba de HeapSort Maximos")
	arreglo := []int{5, 2, 7, 1, 3}
	TDAHeap.HeapSort(arreglo, compararMaximos)

	require.Equal(t, []int{1, 2, 3, 5, 7}, arreglo)

}

func TestHeapSortMin(t *testing.T) {
	t.Log("Prueba de HeapSort Minimos")
	arreglo := []int{5, 2, 7, 1, 3}
	TDAHeap.HeapSort(arreglo, compararMinimos)

	require.Equal(t, []int{7, 5, 3, 2, 1}, arreglo)
}

func TestHeapSortStrings(t *testing.T) {
	t.Log("Prueba de HeapSort con strings")
	arreglo := []string{"casa", "perro", "avion", "zebra", "barco"}
	TDAHeap.HeapSort(arreglo, compararStrings)

	require.Equal(t, []string{"avion", "barco", "casa", "perro", "zebra"}, arreglo)
}

func TestHeapSortPersonasEdad(t *testing.T) {
	t.Log("Prueba de HeapSort con estructuras Persona ordenadas por edad (de el más grande al más chico)")
	personas := []Persona{
		{"Maria", 30},
		{"Pedro", 20},
		{"Ana", 35},
		{"Luis", 28},
	}

	TDAHeap.HeapSort(personas, compararPersonasEdadMax)

	require.Equal(t, []Persona{{"Pedro", 20}, {"Luis", 28}, {"Maria", 30}, {"Ana", 35}}, personas)
}

func TestHeapSortPersonasNombre(t *testing.T) {
	t.Log("Prueba de HeapSort con estructuras Persona ordenadas alfabéticamente por nombre")
	personas := []Persona{
		{"Maria", 30},
		{"Pedro", 20},
		{"Ana", 35},
		{"Luis", 28},
	}

	TDAHeap.HeapSort(personas, compararPersonasNombreMax)

	require.Equal(t, []Persona{{"Ana", 35}, {"Luis", 28}, {"Maria", 30}, {"Pedro", 20}}, personas)
}

func TestVolumenHeapMin(t *testing.T) {
	t.Log("Prueba de volumen heap de mínimos")
	heap := TDAHeap.CrearHeap(compararMinimos)

	for i := 10000; i > 0; i-- {
		heap.Encolar(i)
		require.Equal(t, i, heap.VerMax())
	}

	for i := 1; i <= 10000; i++ {
		require.Equal(t, i, heap.VerMax())
		require.Equal(t, i, heap.Desencolar())
	}

	testHeapVacio(t, heap)
}

func TestVolumenHeapMax(t *testing.T) {
	t.Log("Prueba de volumen heap de máximos")
	heap := TDAHeap.CrearHeap(compararMaximos)

	for i := 0; i < 10000; i++ {
		heap.Encolar(i)
		require.Equal(t, i, heap.VerMax())
	}

	for i := 9999; i >= 0; i-- {
		require.Equal(t, i, heap.VerMax())
		require.Equal(t, i, heap.Desencolar())
	}

	testHeapVacio(t, heap)
}

func BenchmarkHeapMin(b *testing.B) {
	b.Log("Prueba de stress del heap de mínimos.")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenHeapMin(b, n)
			}
		})
	}
}

func BenchmarkHeapMax(b *testing.B) {
	b.Log("Prueba de stress del heap de máximos.")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenHeapMax(b, n)
			}
		})
	}
}

func ejecutarPruebaVolumenHeapMax(b *testing.B, tam int) {
	heap := TDAHeap.CrearHeap(compararMaximos)
	for i := 0; i < tam; i++ {
		heap.Encolar(i)
		require.EqualValues(b, i, heap.VerMax(), "El elemento con máxima prioridad es incorrecto")
	}

	require.EqualValues(b, tam, heap.Cantidad(), "La cantidad de elementos es incorrecta")

	for i := tam - 1; i >= 0; i-- {
		require.Equal(b, i, heap.Desencolar())
	}
	testHeapVacioBenchmark(b, heap)
}

func ejecutarPruebaVolumenHeapMin(b *testing.B, tam int) {
	heap := TDAHeap.CrearHeap(compararMinimos)
	for i := 0; i < tam; i++ {
		heap.Encolar(i)
	}

	require.EqualValues(b, tam, heap.Cantidad(), "La cantidad de elementos es incorrecta")

	for i := 0; i < tam; i++ {
		require.EqualValues(b, i, heap.VerMax(), "El elemento con mínima prioridad es incorrecto")
		require.Equal(b, i, heap.Desencolar())
	}
	testHeapVacioBenchmark(b, heap)
}

func testHeapVacioBenchmark(b *testing.B, heap TDAHeap.ColaPrioridad[int]) {
	require.True(b, heap.EstaVacia())
	require.Equal(b, 0, heap.Cantidad())
	require.PanicsWithValue(b, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(b, "La cola esta vacia", func() { heap.Desencolar() })
}

func BenchmarkHeapSortMax(b *testing.B) {
	b.Log("Prueba de volumen de HeapSort Maximos")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			arreglo := make([]int, n)
			for i := 0; i < n; i++ {
				arreglo[i] = rand.Intn(n * 2)
			}

			TDAHeap.HeapSort(arreglo, compararMaximos)

			for i := 1; i < n; i++ {
				require.True(b, arreglo[i] >= arreglo[i-1])
			}
		})
	}
}

func BenchmarkHeapSortMin(b *testing.B) {
	b.Log("Prueba de volumen de HeapSort Minimos")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			arreglo := make([]int, n)
			for i := 0; i < n; i++ {
				arreglo[i] = rand.Intn(n * 2)
			}

			TDAHeap.HeapSort(arreglo, compararMinimos)

			for i := 1; i < n; i++ {
				require.True(b, arreglo[i] <= arreglo[i-1])
			}
		})
	}
}
