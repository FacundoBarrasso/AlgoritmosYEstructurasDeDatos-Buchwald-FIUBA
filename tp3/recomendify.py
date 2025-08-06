#!/usr/bin/python3

import sys
from auxiliares import parsear_y_construir_bipartito, procesar_stdin

def recomendify():
    grafo = parsear_y_construir_bipartito(sys.argv[1])
    if grafo is not None:
        procesar_stdin(grafo)

if __name__ == "__main__":
    recomendify()