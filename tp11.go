package main

import (
	"fmt"
)

const (
	cv int = 5
	cd int = 4
)

/*
a. Leer y guardar en el vector ART[20] los nombres de los artistas. La posición que ocupe
dentro del vector determinará el código de cada uno, teniendo en cuenta que se enumeran
del 1 al 20.
*/
func main() {
	var (
		ART     [cv]string
		artista string
		//codA    int
		DAT [][cd]string
	)

	for c := 0; c < cv; c++ {
		fmt.Print("Ingrese el nombre del artista: ")
		fmt.Scan(&artista)
		ART[c] = artista
	}

	DAT = albunes(DAT)

	fmt.Print(DAT)
}

func albunes(DAT [][cd]string) [][cd]string {
	var (
		codA, codAl, generoA, año int
	)

		for{
			

		}
}

