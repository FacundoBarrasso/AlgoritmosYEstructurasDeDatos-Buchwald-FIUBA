from grafo import Grafo
from biblioteca import Biblioteca_grafo

def imprimir_grafo(grafo):
    print("Grafo:")
    print(grafo)
    print("-" * 30)

def main():
    # Crear y mostrar los grafos
    grafo1 = Grafo(es_dirigido=False, vertices=["A", "B", "C", "D", "E", "F", "G", "H", "I", "J"])
    grafo1.agregar_arista("A","B")
    grafo1.agregar_arista("A","C")
    grafo1.agregar_arista("A","D")
    grafo1.agregar_arista("B","E")
    grafo1.agregar_arista("C","F")
    grafo1.agregar_arista("D","G")
    grafo1.agregar_arista("B","F")
    grafo1.agregar_arista("H", "I")
    grafo1.agregar_arista("H", "J")
    
    g1 = Biblioteca_grafo(grafo1)
    print(g1.ciclo_largo_n("A", 3))
    print(g1.proyeccion)
    
main()