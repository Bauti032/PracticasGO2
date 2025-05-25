package main

import (
	"fmt"
)

func main() {
	var (
		repeticiones int = 8
		acumulador   int = 0
		N1           int
	)
	for x := 1; x <= repeticiones; x++ {
		fmt.Println("Ingrese el numero: ")
		fmt.Scanln(&N1)
		acumulador += N1

	}
	promedio := acumulador / repeticiones

	fmt.Printf("El promedio es %d  ", promedio)

}
