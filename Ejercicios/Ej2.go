package ejercicio2

import "fmt"

func main() {
	var R float64

	fmt.Print("Ingrese un radio para obtener su longitud y area de circunferencia: ")
	fmt.Scanln(&R)

	longitud := 2 * 3.1416 * R
	Area := R * R * 3.1416

	fmt.Printf("La longitud de la circunferencia es %.2f y el area es %.2f", longitud, Area)
}
