# Pregunta 4

## Para ejecutar

* Instalar Golang versión 1.21.4 o posterior.
* Para compilar y generar un binario ejecutar `go build -o [<nombre_binario>] recursion.go`
* Para ejectuar directamente ingresar `go run recursion.go`

## Consideraciones

* El cliente que ejecuta los benchmarks está programado para siempre almacenar los resultados en un 
archivo de nombre `results.csv`. En caso de querer almacenar varios resultados de mediciones deberá
respaldar cualquier archivo de resultados antes de correr el programa nuevamente. De lo contrario el 
archivo de resultados existente será sobreescrito.
