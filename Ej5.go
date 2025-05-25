package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64

	fmt.Print("Ingrese un numero: ")
	fmt.Scanln(&num)

	R := math.Pow(-1, num)

	if R > 0 {
		fmt.Println("Tu numero es par.")
	} else {
		fmt.Println("Tu numero es impar.")
	}
}
