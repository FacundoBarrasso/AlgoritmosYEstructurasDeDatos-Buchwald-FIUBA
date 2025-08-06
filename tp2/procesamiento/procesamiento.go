package Procesamiento

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	aux "tp2/auxiliares"
)

const (
	LARGO_LEER_ARCHIVO   = 2
	LARGO_VER_VISITANTES = 3
	LARGO_MAS_VISITADOS  = 2
)

func Procesamiento() {
	detector := aux.CrearDetector()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		entrada := strings.Fields(strings.TrimSpace(scanner.Text()))
		if !elegirComando(detector, entrada) {
			continue
		}
	}
	err := scanner.Err()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR %v\n", err)
	}
}

func elegirComando(detector aux.Detector, entrada []string) bool {
	var err error
	switch entrada[0] {
	case "agregar_archivo":
		if errorLargoEntrada(entrada, LARGO_LEER_ARCHIVO) {
			return false
		}
		err := detector.AgregarArchivo(entrada[1])
		if errorComando(entrada[0], err) {
			return false
		}
		return true

	case "ver_visitantes":
		if errorLargoEntrada(entrada, LARGO_VER_VISITANTES) {
			return false
		}
		err = detector.VerVisitantes(entrada[1], entrada[2])
		if errorComando(entrada[0], err) {
			return false
		}
		return true

	case "ver_mas_visitados":
		if errorLargoEntrada(entrada, LARGO_MAS_VISITADOS) {
			return false
		}
		err = detector.MasVisitados(entrada[1])
		if errorComando(entrada[0], err) {
			return false
		}
		return true

	default:
		fmt.Fprintf(os.Stderr, "Error en comando %v\n", entrada[0])
		return false

	}
}

func errorComando(comando string, err error) bool {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error en comando %v\n", comando)
		return true
	}
	return false
}

func errorLargoEntrada(entrada []string, largo int) bool {
	if len(entrada) != largo {
		fmt.Fprintf(os.Stderr, "Error en comando %v\n", entrada[0])
		return true
	}
	return false
}
