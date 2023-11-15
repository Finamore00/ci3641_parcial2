type Arbol<'a> =
    | Hoja of data: 'a
    | Rama of data: 'a * left: Arbol<'a> * right: Arbol<'a>

let rec postOrder (tree: Arbol<'a>): 'a list =
    match tree with
        | Hoja data -> [data]
        | Rama (data, left, right) -> postOrder(left) @ postOrder(right) @ [data]

let rec preOrder (tree: Arbol<'a>): 'a list =
    match tree with
        | Hoja data -> [data]
        | Rama (data, left, right) -> [data] @ preOrder(left) @ preOrder (right)

let rec isMaxHeap (tree: Arbol<'a>): bool = 
    match tree with
        | Hoja _ -> true
        | Rama (data, l, r) -> match l, r with
            | (Hoja datal, Hoja datar) -> data > datal && data > datar
            | (Hoja datal, Rama(datar, rl, rr)) -> data > datal && data > datar && isMaxHeap(Rama(datar, rl, rr))
            | (Rama(datal, ll, lr), Hoja datar) -> data > datar && data > datal && isMaxHeap(Rama(datal, ll, lr))
            | (Rama(datal, ll, lr), Rama(datar, rl, rr)) -> data > datal && data > datar && isMaxHeap(Rama(datal, ll, lr)) && isMaxHeap(Rama(datar, rl, rr))

let isSymmetricalMaxHeap (tree: Arbol<'a>): bool = (preOrder(tree) = postOrder(tree)) && isMaxHeap(tree)