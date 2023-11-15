type Church =
    | Cero
    | Suc of n: Church

let rec suma (a: Church, b: Church): Church =
    match a, b with
        |(Cero, k) -> k
        |(k, Cero) -> k
        |(Suc(n), k) -> suma(n, Suc(k))

let rec multiplicacion (a: Church, b:Church): Church =
    match a, b with
        |(Cero, k) -> Cero
        |(k, Cero) -> Cero
        |(Suc(n), k) -> suma(k, multiplicacion(n, k))