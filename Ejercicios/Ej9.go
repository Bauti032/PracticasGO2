package main

import "fmt"

func main() {
	var (
		cant int
		PU   float64
		PN   float64
	)

	fmt.Print("Ingrese cantidad del producto: ")
	fmt.Scanln(&cant)
	fmt.Print("Ingrese precio unitario del producto: ")
	fmt.Scanln(&PU)

	PN = (float64(cant) * PU) * 1.21

	fmt.Printf("El precio neto es: %.2f", PN)
}
