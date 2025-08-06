package auxiliares

type Detector interface {

	// AgregarArchivo procesa el archivo y detecta posibles casos de ataques de denegación de servicio. En caso de detectar a una IP como sospechosa, imprime por stdout DoS: <IP>
	AgregarArchivo(archivo string) error

	// VerVisitantes imprime de forma ordenada por stdout todas las IPs que realizaron alguna petición dentro de un rango dado, con los límites inclusive
	VerVisitantes(desde, hasta string) error

	// MasVisitados imprime por stdout los n recursos más solicitados ordenados en forma decreciente por cantidad de solicitudes
	MasVisitados(num string) error
}
