# Pregunta 5

Implementación en Golang de adminsitrador de tipos de datos. Lenguajes de Programación I (CI3641), Parcial II,
Pregunta 5. Trimestre Sep-Dic 2023.

## Para ejecutar
* Instalar Go versión 1.21.4 o posterior.
* Para compilar y generar un binario, ejecutar `go build -o <nombre_binario> main/main.go`
* Para correr el programa directamente, ejecutar `go run main/main.go`

## Consideraciones
* Se desconce si el programa funciona con versiones de Go inferiores a la 1.21.4. En este sentido se insta a 
ejecutar el programa con esta versión de Go en pro de evitar inconvenientes
* En caso de desear compilar a un binario, asegurarse de que el nombre del mismo sea distinto a cualquiera de 
los directorios ya existentes (clientFunctions, dataTypes, main). Esto debido a que si se usa cualquiera de estos
nombres se presentan colisiones y el compilador no generará el ejecutable.

## Pruebas unitarias
Para ejecutar las pruebas unitarias creadas para el programa ejecutar el comando `go test ./clientFunctions`