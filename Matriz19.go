// for f1 := 0; f1 < fin-1; f1++ {
// 	for f2:= f1+1; f2 < fin; f2++ {
// 		if c[f1][7] > c[f2][7] {
// 			for c:= 0; c < 8; c++ {
// 				aux := c[f1][c]
// 				c[f1][c] = c[f2][c]
// 				c[f2][c] = aux
// 			}
// 		}
// 	  }
// }
package main

import (
	"fmt"
	// "math/rand"
	// "time"
)

const (
	FilasA    = 3
	ColumnasA = 3

	ColumnasC = 4
	// columnan = 7
)

func main() {

	var (
		MatrizA [FilasA][ColumnasA]int
		// MatrizM [][ColumnasA] int
		// MatrizN [][ColumnasA]int
	)
	// rand.Seed(time.Now().UnixNano())
	matriz := cargamatrizA(MatrizA)
	fmt.Println(matriz)

	matrizm := cargamatrizM(matriz)
	fmt.Println(matrizm)

	matrizc := cargamatrizC(matriz)
	fmt.Println(matrizc)
	//fmt.Println(matriz2)

}

func cargamatrizA(MatrizA [FilasA][ColumnasA]int) [FilasA][ColumnasA]int {

	var nlegajo int
	var nota int

	for i := 0; i < FilasA; i++ {
		fmt.Println("Ingrese el numero de legajo: ")
		fmt.Scan(&nlegajo)
		MatrizA[i][0] = nlegajo
		for j := 1; j < ColumnasA; j++ {
			fmt.Println("Ingrese la nota del tipo: ")
			fmt.Scanln(&nota)
			MatrizA[i][j] = nota //rand.Intn(10)
		}
	}
	return MatrizA
}

func cargamatrizM(MatrizA [FilasA][ColumnasA]int) [][ColumnasA]int{
	var MatrizM [][ColumnasA]int
	// var filaB int
	for i := 0; i < FilasA; i++ {
		contam := 0
		for j :=1; j < ColumnasA; j++ {
			if MatrizA[i][j] > 5 {
				contam ++
			}
		}
		if contam == ColumnasA-1{
			MatrizM = append(MatrizM, MatrizA[i])
		}

	} 
	return MatrizM	//return MatrizM
}


func cargamatrizC(MatrizA [FilasA][ColumnasA]int) [FilasA][ColumnasC]int{
	var matrizC [FilasA][ColumnasC]int
	var acum int
	var prom int

	for i := 0; i < FilasA; i++ {
		for j := 1; j < ColumnasA; j++ {
			acum =+ MatrizA[i][j]
		}
		prom = acum / ColumnasA-1 
		matrizC[i][ColumnasC] = prom
	}
    
	return matrizC

}