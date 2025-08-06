package auxiliares

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	ej "tp0/ejercicios"
)

func LeerArchivo(ruta string) []int {
	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Printf("Error %v al abrir el archivo %s", err, ruta)
		return nil
	}
	defer archivo.Close()

	s := bufio.NewScanner(archivo)
	arreglo := make([]int, 0)
	for s.Scan() {
		num, _ := strconv.Atoi(s.Text())
		arreglo = append(arreglo, num)
	}
	return arreglo
}

func ElegirArchivoMayor(arreglo1, arreglo2 []int) []int {
	if arreglo1 == nil || arreglo2 == nil {
		return nil
	}

	resultado := ej.Comparar(arreglo1, arreglo2)

	var ganador []int

	if resultado >= 0 {
		ganador = arreglo1
	} else {
		ganador = arreglo2
	}

	ej.Seleccion(ganador)

	return ganador
}
