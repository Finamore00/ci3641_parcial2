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