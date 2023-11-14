package calcFunctions

import (
	"errors"
	"strconv"
	"strings"
)

type OpType int8

const (
	PRE  OpType = 0
	POST OpType = 1
	ERR  OpType = -1
)

/*
Función que parsea y verifica la integridad del comando de evaluación.
En caos de que el comando esté correctamente formatteado retorna el string
conteniendo la expresión a ser evaluada y un valor OpType representando
el tipo de expresión que es (Pre-fija o Post-fija). En caso de que haya
algún error con el formateo del comando o con la expresión, se retorna un
valor de error.
*/
func ParseEvalExpr(expr string) ([]string, OpType, error) {
	words := strings.Split(expr, " ")
	//Retornos en caso de error
	errVals := struct {
		expr   []string
		opType OpType
		err    error
	}{
		[]string{},
		ERR,
		errors.New("error parsing eval cmd"),
	}

	if len(words) < 3 {
		return errVals.expr, errVals.opType, errVals.err
	}

	cmdType := strings.ToLower(words[1])

	//Por ahora asumire que las expresiones ingresadas son correctas xd
	if cmdType == "pre" {
		return words[2:], PRE, nil
	}

	if cmdType == "post" {
		return words[2:], POST, nil
	}

	return errVals.expr, errVals.opType, errVals.err
}

/*
Función que parsea y verifica la integridad del comando de evaluación.
En caos de que el comando esté correctamente formatteado retorna el string
conteniendo la expresión a ser evaluada y un valor OpType representando
el tipo de expresión que es (Pre-fija o Post-fija). En caso de que haya
algún error con el formateo del comando o con la expresión, se retorna un
valor de error.
*/
func ParseShowExpr(input string) ([]string, OpType, error) {
	words := strings.Split(input, " ")
	//Retornos en caso de error
	errVals := struct {
		expr []string
		op   OpType
		err  error
	}{
		[]string{},
		ERR,
		errors.New("error parsing show cmd"),
	}

	//Si no se proveen suficientes argumentos
	if len(words) < 3 {
		return errVals.expr, errVals.op, errVals.err
	}

	//Se ve el tipo de operacion ingresada
	cmdType := strings.ToLower(words[1])

	if cmdType == "pre" {
		return words[2:], PRE, nil
	}

	if cmdType == "post" {
		return words[2:], POST, nil
	}

	return errVals.expr, errVals.op, errVals.err

}

/*
Función que recibe un arreglo con los símbolos de una expresión y un indicador
del tipo de expresión que es, y retorna el valor de la expresión evaluada. Se
asume que la expresión ingresada es correcta bajo el tipo de expresión indicada.
*/
func CalculateExprResult(expr []string, op OpType) int {
	operators := []string{"+", "-", "*", "/"}
	stack := []int{}
	var symArr []string
	if op == PRE {
		symArr = reverseSlice(expr)
	} else {
		symArr = expr
	}

	for _, symbol := range symArr {
		//Case read symbol is an operator
		if contains(operators, symbol) {
			op1 := stack[len(stack)-2]
			op2 := stack[len(stack)-1]
			//"Pop" operands
			stack = stack[:len(stack)-2]
			//Perform operation according to operator
			switch symbol {
			case "+":
				stack = append(stack, op1+op2)
			case "-":
				if op == POST {
					stack = append(stack, op1-op2) //Si me da chance quito estos ifs feos
				} else {
					stack = append(stack, op2-op1)
				}
			case "*":
				stack = append(stack, op1*op2)
			case "/":
				if op == POST {
					stack = append(stack, op1/op2)
				} else {
					stack = append(stack, op2/op1)
				}
			}
		} else {
			//If read symbol is a number just push to the stack
			symVal, _ := strconv.Atoi(symbol)
			stack = append(stack, symVal)
		}
	}
	return stack[len(stack)-1] //Result is at top of stack
}

/*
Struct auxiliar para el proceso de conversión de expresiones
a in-fijo
*/
type Expression struct {
	body     string
	operator string
}

/*
Función que recibe una expresión y un indicador del tipo de expresión que es
y retorna un string con la expresión in-fija equivalente.
*/
func ConvertToInfix(exp []string, t OpType) string {
	operators := []string{"+", "-", "*", "/"}
	var symArr []string
	stack := []Expression{}
	var pl, pr int

	if t == PRE {
		symArr = reverseSlice(exp)
		pl = 1
		pr = 2
	} else {
		symArr = make([]string, len(exp))
		copy(symArr, exp)
		pl = 2
		pr = 1
	}

	for _, sym := range symArr {
		//Caso en que se encuentre a un operador
		if contains(operators, sym) {
			//"Pop" a los operandos
			opl := stack[len(stack)-pl]
			opr := stack[len(stack)-pr]
			stack = stack[:len(stack)-2]
			//Se actúa dependiendo del operador
			switch sym {
			case "-":
				if opr.operator == "+" || opr.operator == "-" {
					opr.body = parenth(opr.body)
				}
			case "*", "/":
				if opr.operator != sym && opr.operator != "" {
					opr.body = parenth(opr.body)
				}
				if opl.operator == "+" || opl.operator == "-" {
					opl.body = parenth(opl.body)
				}
			}
			newExpr := Expression{
				body:     opl.body + sym + opr.body,
				operator: sym,
			}
			stack = append(stack, newExpr)
		} else {
			//Caso en que se encuentra un numero
			newExpr := Expression{
				body:     sym,
				operator: "",
			}
			stack = append(stack, newExpr)
		}
	}
	return stack[len(stack)-1].body
}

// Funcion que parentiza una expresion
func parenth(body string) string {
	return "(" + body + ")"
}

/*
Funcion que indica si un slice de strings contiene un string ingresado
porque aparentemente Golang no tiene eso por defecto :)
*/
func contains(sl []string, s string) bool {
	for _, elem := range sl {
		if elem == s {
			return true
		}
	}
	return false
}

/*
Función que recibe un slice de strings sl y retorna un slice que es el resultado
de invertir sl porque aparentemente Golang tampoco tiene eso por defecto :)
*/
func reverseSlice(sl []string) []string {
	l := len(sl)
	newSl := make([]string, l)
	for i := 0; i < l; i++ {
		newSl[l-(i+1)] = sl[i]
	}
	return newSl
}

/*
Función que recibe un slice de strings con los simbolos de una expresion
aritmetica y determina si es una expresion pre-fija valida.
*/
// func isValidPrefix(expr []string) bool {
// 	return true

// }

// func isValidPostfix(expr []string) bool {
// 	return true
// }
