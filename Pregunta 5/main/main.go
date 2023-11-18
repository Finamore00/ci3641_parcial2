/*
Implementación en Golang de simulador de manejador de tipos de datos.
Lenguajes de Programación I (CI3641), Parcial II, Pregunta 5

Autor: Santiago Finamore
Prof. Ricardo Monascal
*/

package main

import (
	"bufio"
	"dataTypeSimulator/clientFunctions"
	"dataTypeSimulator/dataTypes"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Declaracion de variables
	dtDir := map[string]dataTypes.DataType{}
	inputReader := bufio.NewReader(os.Stdin)
	exitFlag := false

	clientFunctions.PrintWelcome()
	for {
		//Solicitar comando del usuario
		fmt.Print("Ingrese el comando a ejecutar: ")
		input, _ := inputReader.ReadString('\n')
		input = strings.TrimSuffix(strings.ToLower(input), "\n") //Remove trailing newline and lowercase

		switch {
		case strings.HasPrefix(input, "atomico "):
			tName, tSize, tAlign, err := clientFunctions.ParseAtomicCmd(input)
			if err != nil {
				clientFunctions.PrintAtomicCmdErr()
				continue
			}
			newType := clientFunctions.NewAtomicType(tName, tSize, tAlign)
			dtDir[tName] = newType
		case strings.HasPrefix(input, "struct "):
			sName, sMembers, err := clientFunctions.ParseStructCmd(input, &dtDir) //Pasamos referencia para no copiar todo el mapa
			if err != nil {
				clientFunctions.PrintStructCmdErr()
				continue
			}
			newType := clientFunctions.NewStructType(sName, sMembers, &dtDir)
			dtDir[sName] = newType
		case strings.HasPrefix(input, "union "):
			uName, uAlts, err := clientFunctions.ParseUnionCmd(input, &dtDir)
			if err != nil {
				clientFunctions.PrintUnionCmdErr()
				continue
			}
			newType := clientFunctions.NewUnionType(uName, uAlts, &dtDir)
			dtDir[uName] = newType
		case strings.HasPrefix(input, "mostrar "):
			tName, err := clientFunctions.ParseShowCmd(input)
			if err != nil {
				clientFunctions.PrintShowCmdErr()
				continue
			}
			ty, ok := dtDir[tName]
			if !ok {
				clientFunctions.PrintTypeNotFound(tName)
				continue
			}
			clientFunctions.ShowDataType(ty)
		case input == "help":
			clientFunctions.PrintHelp()
		case input == "salir":
			exitFlag = true
		default:
			clientFunctions.PrintGenericCmdErr()
		}
		if exitFlag {
			break
		}
	}

}
