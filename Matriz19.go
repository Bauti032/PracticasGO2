package main

import (
	"fmt"
	// "math/rand"
	// "time"
)

const (
	FilasA    = 3
	ColumnasA = 2
	// columnan = 7
)

func main() {

	var (
		MatrizA [FilasA][ColumnasA]int
		MatrizN [][ColumnasA]int
	)
	// rand.Seed(time.Now().UnixNano())
	matriz := cargamatrizA(MatrizA)
	fmt.Println(matriz)

	matriz_2 := cargamatrizN(MatrizN, MatrizA)
	fmt.Println(matriz_2)

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

func cargamatrizN(MatrizN [][ColumnasA]int, MatrizA [FilasA][ColumnasA]int) [][ColumnasA]int {
	var (
		nota	int
		canta	int
	
	)

	for i := 0; i < FilasA; i++ {
		for j := 1; j < ColumnasA; j++ {
			nota = MatrizA[i][j]
			if nota >= 6 {
				canta++
			}
		}
			if canta == 2 {
				for h := 0; h < 2; h ++ {
					MatrizN = append(MatrizN, nota)

				}
			}
	}

	return MatrizN
}
