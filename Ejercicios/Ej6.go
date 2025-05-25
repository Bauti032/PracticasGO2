package main

import "fmt"

func main() {
	var (
		N1 int
		N2 int
		N3 int
	)

	fmt.Println("Ingrese el primer Numero Entero: ")
	fmt.Scanln(&N1)
	fmt.Println("Ingrese el segundo Numero Entero: ")
	fmt.Scanln(&N2)
	fmt.Println("Ingrese el tercer Numero Entero: ")
	fmt.Scanln(&N3)

	if N1 > N2 {
		if N1 > N3 {
			fmt.Printf("El numero %d es mayor", N1)
		} else {
			fmt.Printf("El numero %d es mayor", N2)
		}

	} else if N2 > N3 {
		fmt.Printf("El numero %d es mayor", N2)
	} else {
		fmt.Printf("El numero %d es mayor", N3)
	}
}
