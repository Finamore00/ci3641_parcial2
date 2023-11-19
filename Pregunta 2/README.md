# Pregunta 2

Implementación en Golang de calculadora de expresiones aritméticas. Lenguajes de Programación I (CI3641), Parcial II, Pregunta 2.
Trimestre Sep-Dic 2023

## Para ejecutar

* Instalar Golang versión 1.21.4 o posterior.
* Para compilar y generar un binario ejecutar `go build -o <nombre_binario> main/main.go`
* Para correr directamente ejecutar `go run main/main.go`

## Consideraciones

* Se desconoce si el código proveído funciona con versiones de Golang anteriores a la 1.21.4. Para garantizar la integridad del programa
se insta a utilizar esta versión
* Si se desea compilar, asegurarse de que `<nombre_binario>` sea distinto del nombre "main". Esto debido a que main ya existe dentro del
directorio y el compilador no va a poder generar el ejecutable.
* A pesar de que la calculadora recibe y computa correctamente los resultados para expresiones pre-fijas y post-fijas, la misma no es capaz
de verificar que una expresión pre-fija o post-fija ingresada es una expresión *correcta*. De esta forma por favor abstenerse de ingresar
expresiones mal formateadas dado que el programa arrojará resultados erróneos.

## Pruebas unitarias
Para ejecutar las pruebas unitarias creadas para el programa, ejecutar el comando `go test ./calcFunctions`