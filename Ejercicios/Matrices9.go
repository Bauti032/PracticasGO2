package main

import "fmt"

const (
	F = 5
	C = 3
)

var A[F][C]int

func main() {

	cargamatriz(A)

	Menor, col, fil := menorelemento(&A)
	fmt.Println("El menor elemento es:", Menor)
	fmt.Println("La posición del menor elemento es:", col, fil)
}

func cargamatriz(A [F][C]int) {
	var n int

	for x := 0; x < F; x++ {
		for y := 0; y < C; y++ {
			fmt.Println("Ingrese el número: ")
			fmt.Scanln(&n)
			A[x][y] = n
		}
	}
}

func menorelemento(A *[F][C]int) (menumero int, Columna int, filas int) {
	var Menumero int
	var columna int
	// var filas int 

	for x := 0; x < F; x++ {
		for y := 0; y < C; y++ {
			if y == 0 && x == 0 {
				Menumero = A[x][y]
				columna = x
				filas = y
			} else if A[x][y] < Menumero {
				Menumero = A[x][y]
				columna = x
				filas = y 
			}
		}
	}
	return Menumero, columna, filas
}
