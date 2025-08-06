from collections import deque
from grafo import Grafo
import random
import heapq

COEFICIENTE_AMORTIGUACION = 0.85
ITERACIONES_PAGERANK = 100
LARGO_RANDOMWALK = 10000

class ErrorCamino(Exception):
    pass  

class ErrorCaminoNoPertenece(Exception):
    pass

class ElemCamino:
    def __init__(self, elemento, dato):
        self.elemento = elemento
        self.dato = dato

class Biblioteca_grafo:
    def __init__(self, grafo :Grafo):
        self.bipartito = grafo 
        self.proyeccion = None
        self.pagerank_vertices = None
    
    def reconstruir_camino_bipartito(self, origen, destino, padres :dict, datos :dict):  
        camino = []
        while destino is not origen:
            camino.append(ElemCamino(destino, datos[destino]))
            destino = padres[destino]
        camino.append(ElemCamino(origen, None))
        camino.reverse()
        return camino
    
    def camino_mas_corto(self, origen, destino):
        """Implementa el algoritmo BFS para encontrar el camino más corto entre dos vertices en el grafo bipartito. Devuelve una lista con los vertices que forman el camino más corto o un mensaje de error en caso de no encontrarlo"""
        
        if origen not in self.bipartito or destino not in self.bipartito or " - " not in origen or " - " not in destino:
            return None, ErrorCaminoNoPertenece()

        visitados = set()
        padres = {}
        datos = {}
        cola = deque([origen])
        visitados.add(origen)
        padres[origen] = None
        datos[origen] = None

        while cola:
            actual = cola.popleft()
            if actual == destino:
                return self.reconstruir_camino_bipartito(origen, destino, padres, datos), None

            for adyacente in self.bipartito.adyacentes(actual):
                if adyacente not in visitados:
                    visitados.add(adyacente)
                    padres[adyacente] = actual
                    datos[adyacente] = self.bipartito.peso_arista(actual, adyacente)
                    cola.append(adyacente)

        return None, ErrorCamino()

    def calcular_pagerank(self):
        """Calcula el PageRank de los vertices en el grafo bipartito."""
        
        pagerank = {v: 1 / len(self.bipartito) for v in self.bipartito}
        for _ in range(ITERACIONES_PAGERANK):
            nuevo_pagerank = {}
            for v in self.bipartito:
                suma = 0
                for w in self.bipartito.adyacentes(v):
                    suma += pagerank[w] / len(self.bipartito.adyacentes(w))
                nuevo_pagerank[v] = (1 - COEFICIENTE_AMORTIGUACION) / len(self.bipartito) + COEFICIENTE_AMORTIGUACION * suma
            pagerank = nuevo_pagerank
        return pagerank

    def n_mas_importantes(self, n):
        """Devuelve los n vertices más importantes según el algoritmo de PageRank."""
        
        if self.pagerank_vertices is None:
            pagerank = self.calcular_pagerank()
            self.pagerank_vertices = [(v, pagerank[v]) for v in self.bipartito if " - " in v]
        return heapq.nlargest(n, self.pagerank_vertices, key=lambda x: x[1]) 
        
    def _dividir_bipartito(self, v_inicial, conjuntos):
        conjuntos[v_inicial] = True
        cola = deque([v_inicial])

        while cola:
            v = cola.popleft()
            for w in self.bipartito.adyacentes(v):
                if w in conjuntos:
                    if conjuntos[v] == conjuntos[w]: return False
                else:
                    conjuntos[w] = not conjuntos[v]
                    cola.append(w)  
        return True
      
    def dividir_bipartito(self) ->dict:
        """Devuelve un diccionario con los vertices del grafo bipartito divididos en dos conjuntos disjuntos {v: bool}. En caso de que el grafo no sea bipartito devuelve None"""
            
        conjuntos = {}
        for v in self.bipartito:
            if v not in conjuntos: 
                if not self._dividir_bipartito(v, conjuntos):
                    return None
        return conjuntos

    def proyectar_bipartito(self, v_inicial):
        """Construye la proyeccion de el grafo bipartito"""
    
        conjuntos = self.dividir_bipartito()
        if conjuntos is None or v_inicial not in conjuntos:
            self.proyeccion = None
            return TypeError 
            
        proyeccion = Grafo(False)
        for v in conjuntos:
            if conjuntos[v] == conjuntos[v_inicial]:
                adyacentes = self.bipartito.adyacentes(v)
                for w in adyacentes: proyeccion.agregar_vertice(w)
                
                for i in range(len(adyacentes) - 1):
                    for j in range(i + 1, len(adyacentes)):
                        proyeccion.agregar_arista(adyacentes[i], adyacentes[j])
        self.proyeccion = proyeccion
        return None
       
    def randomwalk(self, origen):
        pagerank = {}
        for j in range(LARGO_RANDOMWALK):
            adyacentes = self.bipartito.adyacentes(origen)
            ady = random.choice(adyacentes)
            pagerank[ady] = pagerank.get(origen, 1) / len(adyacentes)
            origen = ady 
        return pagerank

    def calcular_pagerank_personalizado(self, vertices_interes):
        pagerank_personalizado = {}
        for i in range(ITERACIONES_PAGERANK):
            origen = random.choice(vertices_interes)
            pagerank_randomwalk = self.randomwalk(origen)
            for v in pagerank_randomwalk:
                pagerank_personalizado[v] = pagerank_personalizado.get(v, 0) + pagerank_randomwalk[v]
        return pagerank_personalizado
    
    def n_rank_personalizado(self, v_interes :list, es_grupo_v :bool, n :int):
        """ Aplica el algoritmo PageRank Personalizado empezando desde los vertices en v_interes. Devuelve una lista con los n elementos con mejor rank. es_grupo_v indica si el grupo de la lista resultado coincide con el grupo de los vertices de interes """
        
        pagerank = self.calcular_pagerank_personalizado(v_interes)
        conjuntos = self.dividir_bipartito()
        if es_grupo_v:
            heap = [(-(pagerank.get(v,float("-inf"))), v) for v in conjuntos if conjuntos[v] == conjuntos[v_interes[0]]]
        else:
            heap = [(-(pagerank.get(v, float("-inf"))), v) for v in conjuntos if conjuntos[v] != conjuntos[v_interes[0]]]
        
        ignorados = {v for v in v_interes}
        res = []
        heapq.heapify(heap)
        i = 0
        while len(heap) > 0 and i < n:
            elem = heapq.heappop(heap)
            if elem[1] in ignorados:
                continue
            if elem[0] == float("inf"):
                break
            res.append(elem[1])
            i += 1
        return res

    def reconstruir_camino_ciclo(self, origen, destino, padres :dict):  
        camino = []
        while destino is not None:
            camino.append(destino)
            destino = padres[destino]
        camino.reverse()
        camino.append(origen)
        return camino
    
    def _ciclo_largo_n(self, v, visitados :dict, padres :dict, contador :int):
        if contador < 1:
            return False, v, v, padres
        
        visitados.add(v)
        for w in self.proyeccion.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                es_n_ciclo, origen, destino, _ = self._ciclo_largo_n(w, visitados, padres, contador - 1)
                if es_n_ciclo:
                    return True, origen, destino, padres
            elif padres[v] != w and contador == 1 and padres[w] == None:
                return True, w, v, padres
    
        visitados.remove(v)
        return False, v, v, padres
    
    def ciclo_largo_n(self, origen, n :int) ->list:
        """Devuelve una lista con los vertices que forman un ciclo de largo n empezando por el origen. En caso de no existir tal ciclo, devuelve una lista vacia"""
        
        if n < 3: return []
        
        if self.proyeccion is None:
            err = self.proyectar_bipartito(self.bipartito.adyacentes(origen)[0])
            if err != None: return []
            
        es_ciclo_n, v, destino, padres = self._ciclo_largo_n(origen, set(), {origen: None}, n)
        if es_ciclo_n:
            return self.reconstruir_camino_ciclo(v, destino, padres)
        return []
    
    def vertices_en_rango(self, vertice, n):
        """Devuelve la cantidad de vertices que se encuentran a exactamente n saltos desde el vertice pasado por parámetro utilizando un recorrido BFS en el grafo proyectado. En caso de que haya ocurrido un error al proyectar, devuelve None."""
        
        if self.proyeccion is None:
            err = self.proyectar_bipartito(self.bipartito.adyacentes(vertice)[0])
            if err != None: return None

        vertices_en_rango = 0
        if vertice not in self.proyeccion:
            return vertices_en_rango

        visitados = set()
        cola = deque([(vertice, 0)])
        visitados.add(vertice)

        while cola:
            actual, distancia = cola.popleft()
            if distancia == n:
                vertices_en_rango += 1
            elif distancia < n:
                for adyacente in self.proyeccion.adyacentes(actual):
                    if adyacente not in visitados:
                        visitados.add(adyacente)
                        cola.append((adyacente, distancia + 1))

        return vertices_en_rango