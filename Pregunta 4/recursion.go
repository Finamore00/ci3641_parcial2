/*
Solución implementada en Golang de la respuesta 4. Segundo
Parcial de Lenguajes de Programación I (CI3641), Sep-Dic 2023

Autor: Santiago Finamore
Prof. Ricardo Monascal

Número de carné: 18-10125
	X: 1
	Y: 2
	Z: 5
	alpha: ((X+Y) mod 5) + 3 = 6
	beta: ((Y+Z) mod 5) + 3 = 5
*/

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

/*
Versión recursiva de la función, siguiendo la definición proveída
*/
func F65Rec(n uint64) uint64 {
	if 0 <= n && n < 30 {
		return n
	} else {
		return F65Rec(n-5) + F65Rec(n-10) + F65Rec(n-15) + F65Rec(n-20) + F65Rec(n-25) + F65Rec(n-30)
	}
}

/*
Version con recursion de cola
*/
func F65TailRec(n uint64) uint64 {
	var aux func(uint64, uint64, []uint64) uint64
	baseCases := make([]uint64, 30, 1000) //Capacidad de 1000, para evitar redesignacion de memoria en runtime
	for i := range baseCases {
		baseCases[i] = uint64(i)
	}

	aux = func(n uint64, i uint64, arr []uint64) uint64 {
		if i == n {
			return arr[0]
		} else {
			return aux(n, i+1, append(arr[1:], arr[25]+arr[20]+arr[15]+arr[10]+arr[5]+arr[0]))
		}
	}

	return aux(n, 0, baseCases)
}

/*
Versión iterativa
*/
func F65Iter(n uint64) uint64 {
	i := uint64(0)
	baseCases := make([]uint64, 30, 1000) //Capacidad de 1000, para evitar redesignacion de memoria en runtime
	for i := range baseCases {
		baseCases[i] = uint64(i)
	}

	for i < n {
		baseCases = append(baseCases[1:], baseCases[25]+baseCases[20]+baseCases[15]+baseCases[10]+baseCases[5]+baseCases[0])
		i++
	}

	return baseCases[0]
}

// Pequeña aplicación de medición
func main() {
	casesWithRec := []uint64{10, 50, 70, 80, 90, 100, 120, 140, 160, 180, 200} //Casos para medir
	casesWithoutRec := []uint64{250, 300, 350, 400, 500}
	var tRec time.Duration
	var tTail time.Duration
	var tIter time.Duration
	resArr := [][]string{} //Guardar resultados

	//Create csv file with results
	f, err := os.Create("results.csv")
	w := csv.NewWriter(f)
	if err != nil {
		fmt.Println("Failed to create csv file.")
		os.Exit(1)
	}

	for _, el := range casesWithRec {
		newRes := make([]string, 1, 4)
		newRes[0] = fmt.Sprint(el)
		fmt.Printf("Caso n=%d\n", el)

		t := time.Now()
		res := F65Rec(el)
		tRec = time.Since(t)
		fmt.Println("\tVersión recursiva:")
		fmt.Println("\t\tResultado:", res)
		fmt.Println("\t\tTiempo:", tRec)
		newRes = append(newRes, tRec.String())

		t = time.Now()
		res = F65TailRec(el)
		tTail = time.Since(t)
		fmt.Println("\tVersión recursiva de cola:")
		fmt.Println("\t\tResultado:", res)
		fmt.Println("\t\tTiempo:", tTail)
		newRes = append(newRes, tTail.String())

		t = time.Now()
		res = F65Iter(el)
		tIter = time.Since(t)
		fmt.Println("\tVersión iterativa:")
		fmt.Println("\t\tResultado:", res)
		fmt.Println("\t\tTiempo:", tIter)
		newRes = append(newRes, tIter.String())

		resArr = append(resArr, newRes)
	}

	fmt.Println("Recursión regular removida por exceso de tiempo.")
	for _, el := range casesWithoutRec {
		newRes := make([]string, 1, 4)
		newRes[0] = fmt.Sprint(el)
		newRes = append(newRes, "N/A") //Resultado recursivo
		fmt.Printf("Caso n=%d\n", el)
		t := time.Now()
		res := F65TailRec(el)
		tTail = time.Since(t)
		fmt.Println("\tVersión recursiva de cola:")
		fmt.Println("\t\tResultado:", res)
		fmt.Println("\t\tTiempo:", tTail)
		newRes = append(newRes, tTail.String())

		t = time.Now()
		res = F65Iter(el)
		tIter = time.Since(t)
		fmt.Println("\tVersión iterativa:")
		fmt.Println("\t\tResultado:", res)
		fmt.Println("\t\tTiempo:", tIter)
		newRes = append(newRes, tIter.String())

		resArr = append(resArr, newRes)
	}

	w.WriteAll(resArr)
}
