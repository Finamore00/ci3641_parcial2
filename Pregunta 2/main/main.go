/**
Implementación en Golang de una calculadora simple de expresiones aritméticas.
Pregunta 2 del segundo parcial de Lenguajes de Programación I (CI3641) Sep-Dic 2023

Autor: Santiago Finamore
Prof. Ricardo Monascal
*/

package main

import (
	"bufio"
	"expressionCalculator/calcFunctions"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Dar la bienvenida al programa
	printWelcome()

	var exitFlag bool = false
	//Se comienza ejecución
	for {
		//Obtener comando del usuario
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Print("Ingrese la acción a realizar: ")
		input, _ := inputReader.ReadString('\n')
		input = strings.ToLower(input[:len(input)-1]) //Slice necesario para remover simbolo \n

		switch {
		case strings.HasPrefix(input, "eval "):
			exp, t, err := calcFunctions.ParseEvalExpr(input)
			if err != nil {
				printEvalCmdErr()
			} else {
				fmt.Println(calcFunctions.CalculateExprResult(exp, t))
			}
		case strings.HasPrefix(input, "mostrar "):
			exp, t, err := calcFunctions.ParseShowExpr(input)
			if err != nil {
				printShowCmdErr()
			} else {
				fmt.Println(calcFunctions.ConvertToInfix(exp, t))
			}
		case input == "help":
			printHelp()
		case input == "salir":
			exitFlag = true
		default:
			printCommandErr()
		}
		if exitFlag {
			break
		}
	}
}

/*
Procedimiento que muestra el mensaje de bienvenida al programa.
*/
func printWelcome() {
	fmt.Println("¡Bienvenido/a a la calculadora de expresiones aritméticas!")
	fmt.Println("Para información sobre los comandos disponibles ingresar 'help'")
}

/*
Procedimiento que muestra al usuario los comandos disponibles y su uso
*/
func printHelp() {
	fmt.Println("Comandos disponibles:")
	fmt.Println("\ti. EVAL <orden> <expr>")
	fmt.Println("\t\tMuestra en pantalla el resultado de evaluar la expresión <expr>, escrita en orden <orden>")
	fmt.Println("\t\tÓrdenes disponibles:")
	fmt.Println("\t\t\t- PRE: Expresiones pre-fijas")
	fmt.Println("\t\t\t- POST: Expresiones post-fijas")
	fmt.Println("\tii. MOSTRAR <orden> <expr>")
	fmt.Println("\t\tMuestra el equivalente in-fijo de la expresión <expr>, que está escrita en el orden <orden>")
	fmt.Println("\tiii. SALIR")
	fmt.Println("\t\tCulmina el programa")
}

/*
Indica al usuario que cometió algún error ingresando un comando
*/

func printCommandErr() {
	fmt.Println("\tComando ingresado descononcido o mal formateado.")
	fmt.Println("\tPara ver comandos disponibles y su uso, ingresar 'help'")
}

/*
Muestra al usuario que cometió un error en el comando de evaluación
*/
func printEvalCmdErr() {
	fmt.Println("\tEl comando de evaluación ingresado no es válido o la expresión ingresada no es conforme con el orden.")
	fmt.Println("\tPara consultar uso del comando EVAL ver 'help'.")
	fmt.Println("\tNOTA: Los símbolos de la expresión deben estar propiamente separados por espacios.")
}

/*
Muestra al usuario que cometió un error en el comando de muestra
*/
func printShowCmdErr() {
	fmt.Println("\t El comando de muestra ingresado no es válido o la expresión ingresada no es conforme con el orden.")
	fmt.Println("\tPara consultar uso del comando MOSTRAR ver 'help'")
	fmt.Println("\tNOTA: Los símbolos de la expresión deben estar propiamente separados por espacios.")
}
