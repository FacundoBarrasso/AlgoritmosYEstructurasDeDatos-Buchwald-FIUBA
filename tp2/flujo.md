## Agregar archivo
El comando se acompaña de la ruta de un archivo de log, accesible desde el mismo directorio donde se ejecuta el programa.

Ejemplo: agregar_archivo 20171025.log

Al ejecutarse se deberá procesar el archivo, y detectar posibles casos de ataques de denegación de servicio. Si se detecta que una dirección IP realizó cinco o más peticiones en menos de dos segundos, el comando debe alertarlo por salida estándar como sospechosa de intento de DoS.

A la hora de detectar denegaciones de servicio, varios archivos se consideran independientes ente sí, por lo que no se deberán memorizar entradas entre dos ejecuciones diferentes de agregar_archivo.

Para alertar una IP, basta con mostrar por salida estándar DoS: <IP>. Una misma dirección no deberá ser reportada más de una vez. Si varias direcciones son sospechosas, estas deberán ser mostradas en orden creciente, numéricamente.

### Complejidad:
* Búsquedas O(N) (N = cant líneas del log)
* Mantenimiento actualizar los más visitados O(N)
* Mantenimiento actualizar los visitantes (N log V) (V = cant visitantes en toda la historia del programa)

### Implementación (Facu):

voy leyendo los logs

llamar a parsear logs --> struct {IP, Fecha hora, Método, URL}

lleno ABB con IP (estructura "global")
lleno hash con URL:cant visitas (estructura "global")


tengo cola que funciona como sliding window 
encolo elementos hasta que dif(tiempo primero, tiempo nuevo elem) > 2seg
voy desencolando hasta que dif < 2 o cola vacia --> cuando desencolo agrego a dic 
    *{dir IP: cant visitas}*
   
cuando encolo elemento itero sobre dic anterior y verifico que ninguno tenga cant > 5 --> si DoS lo encolo en estructura DoS

estructura DoS = arreglo 
    `
        Dir IP tipo de dato struct de 4 bloques numéricos (c/u 1 byte) --> ordeno con Radix
    `

devuelve estructura DoS = arreglo

**Pili**: funcion general llama a funcion de facu, ordena el arreglo con radix y devuelve por stdout

## Ver visitantes (Facu)
ABB con IPs --> iterador por rango
func cmp radix 

## Ver más visitados (Pili)
agarra hash y lo mete al heap
Heap --> algoritmo top K