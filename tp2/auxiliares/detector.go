package auxiliares

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	TDACola "tdas/cola"
	TDAHeap "tdas/cola_prioridad"
	TDADic "tdas/diccionario"
	"time"
)

const (
	CANTIDAD_IP_DOS  = 5
	DISTANCIA_TIEMPO = 2 * time.Second
)

type detector struct {
	ips         TDADic.DiccionarioOrdenado[ip, int]
	accesosURLs TDADic.Diccionario[string, int]
}

func CrearDetector() Detector {
	return &detector{
		ips:         TDADic.CrearABB[ip, int](comp),
		accesosURLs: TDADic.CrearHash[string, int](),
	}
}

func (detector *detector) AgregarArchivo(archivo string) error {
	dos, err := detector.agregarArchivo(archivo)
	if err != nil {
		return err
	}
	imprimirDos(dos)
	fmt.Println("OK")
	return nil
}

func (detector *detector) agregarArchivo(rutaLogs string) ([]ip, error) {
	dicDoS, err := detector.procesarArchivo(rutaLogs)
	if err != nil {
		return nil, err
	}

	arrDoS := make([]ip, 0, dicDoS.Cantidad())
	for iter := dicDoS.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		ip, _ := iter.VerActual()
		arrDoS = append(arrDoS, ip)
	}

	dosOrdenado := ordernarIps(arrDoS)

	return dosOrdenado, nil
}

func (detector *detector) procesarArchivo(rutaLogs string) (TDADic.Diccionario[ip, int], error) {
	archivo, err := os.Open(rutaLogs)
	if err != nil {
		err := fmt.Errorf("error al abrir la línea de log: %s", err)
		return nil, err
	}
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)

	ipsDos, err := detector.detectarDoS(scanner)
	if err != nil {
		return nil, err
	}

	if err := scanner.Err(); err != nil {
		err := fmt.Errorf("error al leer el archivo: %s", err)
		return nil, err
	}

	return ipsDos, err
}

func (detector *detector) detectarDoS(scanner *bufio.Scanner) (TDADic.Diccionario[ip, int], error) {
	ipsDos := TDADic.CrearHash[ip, int]()
	slidingWindow := TDACola.CrearColaEnlazada[entradaLog]()
	ipsEnCola := TDADic.CrearHash[ip, int]()

	for scanner.Scan() {
		linea := scanner.Text()
		log, err := parsearLinea(linea)
		if err != nil {
			return nil, err
		}

		slidingWindow.Encolar(log)
		if !ipsEnCola.Pertenece(log.ip) {
			ipsEnCola.Guardar(log.ip, 1)
		} else {
			ipsEnCola.Guardar(log.ip, ipsEnCola.Obtener(log.ip)+1)
		}

		if !(detector.accesosURLs).Pertenece(log.url) {
			(detector.accesosURLs).Guardar(log.url, 1)
		} else {
			(detector.accesosURLs).Guardar(log.url, (detector.accesosURLs).Obtener(log.url)+1)
		}
		(detector.ips).Guardar(log.ip, 0)

		for !slidingWindow.EstaVacia() && log.momento.Sub(slidingWindow.VerPrimero().momento) >= DISTANCIA_TIEMPO {
			primero := slidingWindow.Desencolar()
			ipsEnCola.Guardar(primero.ip, ipsEnCola.Obtener(primero.ip)-1)
		}

		if ipsEnCola.Obtener(log.ip) >= CANTIDAD_IP_DOS && !ipsDos.Pertenece(log.ip) {
			ipsDos.Guardar(log.ip, 0)
		}
	}
	return ipsDos, nil
}

func (detector *detector) VerVisitantes(desde, hasta string) error {
	visitantes, err := detector.verVisitantes(desde, hasta)
	if err != nil {
		return err
	}
	imprimirVisitantes(visitantes)
	fmt.Println("OK")
	return nil
}

func (detector *detector) verVisitantes(desde, hasta string) ([]ip, error) {
	desdeIp, err := separarIP(desde)
	if err != nil {
		return nil, err
	}

	hastaIp, err := separarIP(hasta)
	if err != nil {
		return nil, err
	}

	visitantes := make([]ip, 0, (detector.ips).Cantidad())

	(detector.ips).IterarRango(&desdeIp, &hastaIp, func(clave ip, dato int) bool {
		visitantes = append(visitantes, clave)
		return true
	})

	return visitantes, nil
}

func (detector *detector) MasVisitados(num string) error {
	k, err := strconv.Atoi(num)
	if err != nil {
		return err
	}
	masVisitados := detector.masVisitados(k)
	imprimirMasVisitados(masVisitados)
	fmt.Println("OK")
	return nil
}

func (detector *detector) masVisitados(k int) []visita {
	arr := make([]visita, 0, (detector.accesosURLs).Cantidad())

	for iter := (detector.accesosURLs).Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		recurso, solicitudes := iter.VerActual()
		arr = append(arr, visita{recurso, solicitudes})
	}

	heap := TDAHeap.CrearHeapArr(arr, func(a, b visita) int {
		return a.cantidadSolicitudes - b.cantidadSolicitudes
	})

	top := make([]visita, 0, k)
	for i := 0; i < k && !heap.EstaVacia(); i++ {
		top = append(top, heap.Desencolar())
	}

	return top
}
