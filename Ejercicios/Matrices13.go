package main 

import (
	"fmt"
	"math/rand"
)

const (
	F = 10
	C = 20
)


func main() {

var (
	A[F][C]int
	Par bool
	Mayor bool
)

	cargarMatriz(&A)

	ordenar(A[F][C], Par, Mayor)


}

func cargarMatriz(A*[F][C]int){

	for f := 0; f< F; f++{
		for c := 0; c<C; c++{
			A[f][c] = rand.Intn(100)
		}
	}


}

func ordenar(A[10][20]int, Par bool, Mayor bool) int {
	var (
		Nmayor int
		Nmenor int
		R string 
		i int
	)
	fmt.Println("Desea buscar por filas pares o impares?: ")
	fmt.Scanln(&R)

	if R == "Par" {
		Par = true
	} else {
		Par = false
	}

	fmt.Println("Desea busca el Numero mayor o menor?: ")
	fmt.Scanln(&R)

	if R == "Mayor" {
		Mayor = true
	} else {
		Mayor = false
	}

	if Par == true {
		i = 0
	} else {
		i = 1
	}

	for i; i < F; i +2 {
		for	j:= 0; j < C; j++ {
			if Mayor == true {
				if A[i][j] > Nmayor {
					Nmayor = A[i][j]
				}
			} else {
				if A[i][j] < Nmenor {
					Nmenor = A[i][j]
				}
			}
		}
		// if Mayor == true{
    	// 	return Nmayor
		// } else {
    	// 	return Nmenor
		// }
	}
}
