package operaciones

import (
	"fmt"
	"math"
)

type operacion struct {
	operador      string
	cantOperandos int
	fn            func(operandos []int64) (int64, error)
}

const (
	OPERANDOS_SUMA           = 2
	OPERANDOS_RESTA          = 2
	OPERANDOS_MULTIPLICACION = 2
	OPERANDOS_DIVISION       = 2
	OPERANDOS_RAIZ           = 1
	OPERANDOS_POTENCIA       = 2
	OPERANDOS_LOGARITMO      = 2
	OPERANDOS_TERNARIO       = 3
)

var operaciones = []operacion{
	// Suma (a + b)
	{"+", OPERANDOS_SUMA, func(operandos []int64) (int64, error) {
		return operandos[1] + operandos[0], nil
	}},

	// Resta (a - b)
	{"-", OPERANDOS_RESTA, func(operandos []int64) (int64, error) {
		return operandos[1] - operandos[0], nil
	}},

	// Multiplicación (a * b)
	{"*", OPERANDOS_MULTIPLICACION, func(operandos []int64) (int64, error) {
		return operandos[1] * operandos[0], nil
	}},

	// División entera (a / b)
	{"/", OPERANDOS_DIVISION, func(operandos []int64) (int64, error) {
		if operandos[0] == 0 {
			return 0, fmt.Errorf("ERROR")
		}
		return operandos[1] / operandos[0], nil
	}},

	// Raiz cuadrada entera (√a)
	{"sqrt", OPERANDOS_RAIZ, func(operandos []int64) (int64, error) {
		if operandos[0] < 0 {
			return 0, fmt.Errorf("ERROR")
		}
		return int64(math.Sqrt(float64(operandos[0]))), nil
	}},

	// Potencia (a ^ b)
	{"^", OPERANDOS_POTENCIA, func(operandos []int64) (int64, error) {
		if operandos[0] < 0 {
			return 0, fmt.Errorf("ERROR")
		}
		return int64(math.Pow(float64(operandos[1]), float64(operandos[0]))), nil
	}},

	// Logaritmo entero (log a b = log b / log a)
	{"log", OPERANDOS_LOGARITMO, func(operandos []int64) (int64, error) {
		if operandos[0] < 2 || operandos[1] <= 0 {
			return 0, fmt.Errorf("ERROR")
		}
		return int64(math.Log(float64(operandos[1])) / math.Log(float64(operandos[0]))), nil
	}},

	// Ternario (condicion ? a : b)
	{"?", OPERANDOS_TERNARIO, func(operandos []int64) (int64, error) {
		if operandos[2] != 0 {
			return operandos[1], nil
		}
		return operandos[0], nil
	}},
}

func (operacion operacion) operar(operandos []int64) (int64, error) {
	if len(operandos) != operacion.cantOperandos {
		return 0, fmt.Errorf("ERROR")
	}
	return operacion.fn(operandos)
}

// Operar recibe un operador y una lista de operandos y devuelve el resultado de la operación. Si el operador no existe devuelve 0 y un error.
func Operar(operador string, operandos []int64) (int64, error) {
	for _, op := range operaciones {
		if op.operador != operador {
			continue
		}
		return op.operar(operandos)
	}
	return 0, fmt.Errorf("ERROR")
}

// CantOperandos devuelve la cantidad de operandos que recibe un operador. Si el operador no existe devuelve -1.
func CantOperandos(operador string) int {
	for _, op := range operaciones {
		if op.operador != operador {
			continue
		}
		return op.cantOperandos
	}
	return -1
}
