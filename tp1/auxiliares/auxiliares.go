package auxiliares

import (
	"bufio"
	Dc "dc/operaciones"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAPila "tdas/pila"
)

const (
	BASE    = 10
	BITSIZE = 64
	ERROR   = "ERROR"
)

func Calculadora() {
	lector := bufio.NewScanner(os.Stdin)
	for lector.Scan() {
		linea := lector.Text()
		if linea == "" {
			break
		}
		resultado := procesarOperacion(linea)
		fmt.Println(resultado)
		if lector.Err() != nil {
			fmt.Println("Error al leer la entrada:", lector.Err())
		}
	}
}

func procesarOperacion(operacion string) string {
	pila := TDAPila.CrearPilaDinamica[int64]()
	arr := strings.Fields(operacion)
	for _, caracter := range arr {
		num, err := strconv.ParseInt(caracter, BASE, BITSIZE)
		if err == nil {
			pila.Apilar(num)
		} else {
			err := procesarOperador(caracter, pila)
			if err != nil {
				return ERROR
			}
		}
	}
	return obtenerResultadoFinal(pila)
}

func procesarOperador(operador string, pila TDAPila.Pila[int64]) error {
	cant := Dc.CantOperandos(operador)
	if cant == -1 {
		return errors.New(ERROR)
	}
	operandos, err := obtenerOperandos(pila, cant)
	if err != nil {
		return err
	}
	res, err := Dc.Operar(operador, operandos)
	if err != nil {
		return err
	}
	pila.Apilar(res)
	return nil
}

func obtenerOperandos(pila TDAPila.Pila[int64], cant int) ([]int64, error) {
	operandos := make([]int64, cant)
	for i := 0; i < cant; i++ {
		if pila.EstaVacia() {
			return nil, errors.New(ERROR)
		}
		operandos[i] = pila.Desapilar()
	}
	return operandos, nil
}

func obtenerResultadoFinal(pila TDAPila.Pila[int64]) string {
	if pila.EstaVacia() {
		return ERROR
	}
	final := pila.Desapilar()
	if !pila.EstaVacia() {
		return ERROR
	}
	return strconv.FormatInt(final, BASE)
}
