package main

import (
	"fmt"
)

const (
	cantcolumnas = 6
	cantfilas    = 7
)

func main() {

	var (
		matriz [cantfilas][cantcolumnas]int
	)

	cargamatriz(&matriz)
	for _, fila := range matriz {
		fmt.Printf("Zona: %d | ", fila[0])
		for j := 1; j < len(fila); j++ {
			fmt.Printf("Prod%d: %d ", j, fila[j])
		}
		fmt.Println()
	}

}

func cargamatriz(matriz *[cantfilas][cantcolumnas]int) [cantfilas][cantcolumnas]int {

	var (
		prod    int
		zona    int
		norden  int
		totzona int
	)

	for {
		fmt.Println("Ingrese el Nro de Orden: ")
		fmt.Scanln(&norden)

		if norden == 0 {
			break
		}

		fmt.Println("Ingrese el Nro de zona comprendido entre 1 y 5: ")
		fmt.Scanln(&zona)

		matriz[zona][0] = zona

		for c := 1; c < cantcolumnas-1; c++ {
			fmt.Printf("Ingrese la cantidad del producto N° %d: ", c)
			fmt.Scanln(&prod)
			matriz[zona][c] = +prod
		}

		for f := 0; f < cantfilas-1; f++ {
			for c := 1; c < cantcolumnas-1; c++ {
				totzona += matriz[f][c]

			}
			matriz[f][cantcolumnas-1] = totzona
		}

	}

	return *matriz
}
