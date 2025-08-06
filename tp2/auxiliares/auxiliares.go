package auxiliares

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	RANGO_IP = 256
	LAYOUT   = "2006-01-02T15:04:05-07:00"
	BASE     = 10
	BITS     = 16
)

type entradaLog struct {
	ip      ip
	momento time.Time
	metodo  string
	url     string
}

type visita struct {
	recurso             string
	cantidadSolicitudes int
}

type ip struct {
	bloque1 int64
	bloque2 int64
	bloque3 int64
	bloque4 int64
}

func parsearLinea(linea string) (entradaLog, error) {
	elementos := strings.Split(linea, "\t")
	if len(elementos) != 4 {
		err := fmt.Errorf("linea de log inválida: %s", linea)
		return entradaLog{}, err
	}
	momento, err := time.Parse(LAYOUT, elementos[1])
	if err != nil {
		return entradaLog{}, err
	}

	ip, err := separarIP(elementos[0])
	if err != nil {
		return entradaLog{}, err
	}

	return entradaLog{
		ip:      ip,
		momento: momento,
		metodo:  elementos[2],
		url:     elementos[3],
	}, nil
}

func separarIP(ipStr string) (ip, error) {
	ipSeparada := strings.Split(ipStr, ".")
	var ipNueva ip
	for i := 0; i < 4; i++ {
		parte, err := strconv.ParseInt(ipSeparada[i], BASE, BITS)
		if err != nil {
			return ip{}, err
		}
		switch i {
		case 0:
			ipNueva.bloque1 = parte
		case 1:
			ipNueva.bloque2 = parte
		case 2:
			ipNueva.bloque3 = parte
		case 3:
			ipNueva.bloque4 = parte
		}
	}
	return ipNueva, nil
}

func countingIps(ips []ip, ipAIndice func(ip) int64) []ip {
	frequencias := make([]int, RANGO_IP)
	for _, dirIp := range ips {
		frequencias[ipAIndice(dirIp)]++
	}

	inicios := make([]int, RANGO_IP)
	for i := 1; i < RANGO_IP; i++ {
		inicios[i] = inicios[i-1] + frequencias[i-1]
	}

	ordenadas := make([]ip, len(ips))
	for _, dirIp := range ips {
		ordenadas[inicios[ipAIndice(dirIp)]] = dirIp
		inicios[ipAIndice(dirIp)]++
	}
	return ordenadas
}

func ordernarIps(ips []ip) []ip {
	ordenado4 := countingIps(ips, func(ip ip) int64 { return ip.bloque4 })
	ordenado3 := countingIps(ordenado4, func(ip ip) int64 { return ip.bloque3 })
	ordenado2 := countingIps(ordenado3, func(ip ip) int64 { return ip.bloque2 })
	return countingIps(ordenado2, func(ip ip) int64 { return ip.bloque1 })
}

func imprimirDos(ipsDos []ip) {
	for _, dirIp := range ipsDos {
		fmt.Printf("DoS: %d.%d.%d.%d\n", dirIp.bloque1, dirIp.bloque2, dirIp.bloque3, dirIp.bloque4)
	}
}

func imprimirVisitantes(visitantes []ip) {
	fmt.Println("Visitantes:")
	for _, dirIp := range visitantes {
		fmt.Printf("\t%d.%d.%d.%d\n", dirIp.bloque1, dirIp.bloque2, dirIp.bloque3, dirIp.bloque4)
	}
}

func imprimirMasVisitados(visitas []visita) {
	fmt.Println("Sitios más visitados:")
	for _, visita := range visitas {
		fmt.Printf("\t%v - %d\n", visita.recurso, visita.cantidadSolicitudes)
	}
}

func comp(a, b ip) int {
	if a.bloque1 != b.bloque1 {
		if a.bloque1 < b.bloque1 {
			return -1
		}
		return 1
	}
	if a.bloque2 != b.bloque2 {
		if a.bloque2 < b.bloque2 {
			return -1
		}
		return 1
	}
	if a.bloque3 != b.bloque3 {
		if a.bloque3 < b.bloque3 {
			return -1
		}
		return 1
	}
	if a.bloque4 != b.bloque4 {
		if a.bloque4 < b.bloque4 {
			return -1
		}
		return 1
	}
	return 0
}
