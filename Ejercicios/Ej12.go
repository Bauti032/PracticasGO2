package main

import (
	"fmt"
)

var (
	N1        int
	Multiplos int
)

func main() {
	for x := 1; x <= 10; x++ {
		fmt.Println("Ingrese el Numero: ")
		fmt.Scanln(&N1)

		if N1%5 == 0 {
			Multiplos += 1
		}
	}
	fmt.Println("La cantidad de multiplos es: ", Multiplos)
}
