from random import choice

class Grafo:
    def __init__(self, es_dirigido, vertices=None):
        self.vertices = {}
        if vertices is not None:
            for v in vertices:
                self.vertices[v] = {}
        self.es_dirigido = es_dirigido

    def __len__ (self):
        return len(self.vertices)

    def __str__(self):
        resultado = []
        if self.es_dirigido:
            for v, ady in self.vertices.items():
                for w, peso in ady.items():
                    resultado.append(f"{v} -> {w} (peso: {peso})")
        else:
            visitados = set()
            for v, ady in self.vertices.items():
                for w, peso in ady.items():
                    if (w,v) not in visitados:
                        visitados.add((v,w))
                        resultado.append(f"{v} -- {w} (peso: {peso})")
        return "\n".join(resultado)
    
    def __contains__(self, v):
        return v in self.vertices
    
    def __iter__(self):
        return iter(self.vertices)

    def agregar_vertice(self, v):
        if v not in self.vertices:
            self.vertices[v] = {}

    def borrar_vertice(self, v):
        if v not in self.vertices:
            raise NameError(f"El vertice {v} no existe")
        self.vertices.pop(v)
        for vertice in self.vertices.keys():
            self.vertices[vertice].pop(v, None)

    def agregar_arista(self, v, w, peso = 1):
        if v not in self.vertices or w not in self.vertices:
            raise NameError("El vertice no existe")
        if v == w:
            raise ValueError(f"{v} no puede tener una arista consigo mismo")
        self.vertices[v][w] = peso
        if not self.es_dirigido:
            self.vertices[w][v] = peso

    def estan_unidos(self, v, w):
        if v not in self.vertices or w not in self.vertices:
            raise NameError("El vertice no existe")
        peso = self.vertices[v].get(w, None)
        return peso is not None

    def borrar_arista(self, v, w):
        if not self.estan_unidos(v, w):
            raise ValueError(f"{v} y {w} no están unidos")
        self.vertices[v].pop(w)
        if not self.es_dirigido:
            self.vertices[w].pop(v)

    def peso_arista(self, v, w):
        if not self.estan_unidos(v, w):
            raise ValueError(f"{v} y {w} no están unidos")
        return self.vertices[v][w]

    def obtener_vertices(self):
        return list(self.vertices.keys())

    def vertice_aleatorio(self):
        if len(self.vertices) == 0:
            raise ValueError("El grafo está vacío")
        return choice(list(self.vertices.keys()))

    def adyacentes(self, v):
        if v not in self.vertices:
            raise NameError(f"El vertice {v} no existe")
        return list(self.vertices[v].keys())
