package main

import "fmt"

func main() {

	var Notas [4][4]int
	var N int

	for f := 0; f < 4; f++ {
		for c := 0; c < 4; c++ {
			fmt.Printf("Ing.Nota %d Alumno %d: ", c+1, f+1)
			fmt.Scanln(&N)
			Notas[f][c] = N
		}
	}

	fmt.Println("Matriz de Notas:")
	for f := 0; f < 4; f++ {
		for c := 0; c < 4; c++ {
			fmt.Printf("%d\t", Notas[f][c])
		}
		fmt.Println()
	}
}
