import sys
from biblioteca import Biblioteca_grafo, ErrorCamino, ErrorCaminoNoPertenece
from grafo import Grafo

LONGITUD_DATOS = 7
MODO_LECTURA = "r"

def procesar_stdin(grafo :Biblioteca_grafo):
    stdin = sys.stdin.read().splitlines()
    for linea in stdin:
        entrada = linea.strip().split(" ", 1)
        
        if entrada[0] == "camino":
            origen, destino = procesar_entrada_camino(entrada)
            camino, err = grafo.camino_mas_corto(origen, destino)
            procesar_salida_camino(camino, err)

        elif entrada[0] == "mas_importantes":
            if not entrada[1].isdigit():
                continue
            resultado = grafo.n_mas_importantes(int(entrada[1]))
            procesar_salida_mas_importantes(resultado)

        elif entrada[0] == "recomendacion":
            canciones, es_cancion, cantidad, err = procesar_entrada_recomendacion(entrada)
            if err is not None:
                continue
            recomendacion = grafo.n_rank_personalizado(canciones, es_cancion, cantidad)
            print("; ".join(recomendacion))

        elif entrada[0] == "ciclo":
            origen, n, err = procesar_entrada_ciclo(entrada)
            if err is not None:
                continue
            ciclo = grafo.ciclo_largo_n(origen, n)
            procesar_salida_ciclo(ciclo)

        elif entrada[0] == "rango":
            cancion, n, err = procesar_entrada_rango(entrada)
            if err is not None:
                continue
            print(grafo.vertices_en_rango(cancion, n))
            
        else:
            continue
                
def parsear_y_construir_bipartito(archivo_tsv):
    grafo = Grafo(es_dirigido=False)
    try:
        with open(archivo_tsv, MODO_LECTURA) as archivo:
            next(archivo)
            for linea in archivo:
                datos = linea.strip().split('\t')
                if len(datos) < LONGITUD_DATOS:
                    continue

                user_id = datos[1]
                track_name = datos[2]
                artist = datos[3]
                playlist_name = datos[5]
                cancion = f"{track_name} - {artist}"

                grafo.agregar_vertice(user_id)
                grafo.agregar_vertice(cancion)
                grafo.agregar_arista(user_id, cancion, playlist_name)
        return Biblioteca_grafo(grafo)
    except FileNotFoundError:
        print("No se encontro el archivo")
        return None
    
def procesar_entrada_camino(entrada):
    parametros = entrada[1].split(" >>>> ")
    return parametros[0], parametros[1]
    
def procesar_salida_camino(camino, err):
    if type(err) == ErrorCamino:
        print("No se encontro recorrido")
    elif type(err) == ErrorCaminoNoPertenece:
        print("Tanto el origen como el destino deben ser canciones")
    else:
        i = 0
        while i < len(camino) - 2:
            print(f"{camino[i].elemento} --> aparece en playlist --> {camino[i + 1].dato} --> de --> {camino[i + 1].elemento} --> tiene una playlist --> {camino[i + 2].dato} --> donde aparece --> ", end = "")
            i += 2
        print(f"{camino[len(camino)-1].elemento}")

def procesar_salida_mas_importantes(vertices):
    for i in range(len(vertices)-1):
        print(f"{vertices[i][0]}; ", end="")
    print(vertices[len(vertices)-1][0])

def procesar_entrada_recomendacion(entrada):
    parametros = entrada[1].split(" ", 2)
    if not parametros[1].isdigit():
        return None, None, None, TypeError
    canciones = parametros[2].split(" >>>> ")
    return canciones, True if parametros[0] == "canciones" else False, int(parametros[1]), None

def procesar_entrada_ciclo(entrada):
    parametros = entrada[1].split(" ", 1)
    if not parametros[0].isdigit():
        return None, None, TypeError
    return parametros[1], int(parametros[0]), None

def procesar_salida_ciclo(ciclo):
    print(" --> ".join(ciclo)) if len(ciclo) != 0 else print("No se encontro recorrido")

def procesar_entrada_rango(entrada):
    parametros = entrada[1].split(" ", 1)
    if not parametros[0].isdigit():
        return None, None, TypeError
    return parametros[1], int(parametros[0]), None