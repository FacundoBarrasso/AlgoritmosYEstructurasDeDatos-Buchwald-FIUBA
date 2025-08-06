package main

import (
	"fmt"
	aux "tp0/auxiliares"
)

const ruta1 = "archivo1.in"
const ruta2 = "archivo2.in"

func main() {

	arreglo1 := aux.LeerArchivo(ruta1)
	arreglo2 := aux.LeerArchivo(ruta2)

	arregloMayor := aux.ElegirArchivoMayor(arreglo1, arreglo2)

	for _, elem := range arregloMayor {
		fmt.Println(elem)
	}
}
