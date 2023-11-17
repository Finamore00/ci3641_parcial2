'''
Implementación en Python de iterador que dada una lista
de enteros positivos no repetidos, devuelve todas las sublistas
de la lista ingresada donde los elementos están dispuestos en orden
creciente. Lenguajes de Programación I (CI3641), Parcial II, Pregunta 3.b

Autor: Santiago Finamore
Prof. Ricardo Monascal
'''

def listas_crecientes(x: list[int]):
    def sub_listas(y: list[int]):
        if y == []:
            yield y
        else:
            for x in sub_listas(y[1:]):
                yield x
                yield [y[0], *x]
    
    for l in sub_listas(x):
        if all(l[i] <= l[i+1] for i in range(len(l)-1)): #ver si la lista está ordenada
            yield l