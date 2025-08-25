package main

import (
	"fmt"
	"math/rand"
)

const (
	F = 10
	C = 20
)
var (
	Mayor bool
	Par bool
	
)

func main() {

	cargamatriz()
	decisiones()
	

}

func cargamatriz() [F][C]int {
	var A[F][C] int

	for i := 0; i < F; i++ {
		for j := 0; j < C; j++ {
			A[i][j] = rand.Intn(100)
		}
	}
	return A
}
func decisiones() {
	var (
		R string
	)

	fmt.Println("Quieres buscar el mayor o menor?: ")
	fmt.Scanln(&R)

	if R == "Mayor" {
		Mayor = true
	} else {
		Mayor = false
	}

	fmt.Println("Quieres buscar por filas pares o impares?: ")
	fmt.Scanln(&R)

	if R == "Pares" {
		Par = true
	} else {
		Par = false
	}
}