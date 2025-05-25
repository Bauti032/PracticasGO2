package main

import "fmt"

func main() {
	var (
		N1 int
		N2 int
		N3 int
	)

	fmt.Print("Ingrese el 1° numero: ")
	fmt.Scanln(&N1)
	fmt.Print("Ingrese el 2° numero: ")
	fmt.Scanln(&N2)
	fmt.Print("Ingrese el 3° numero: ")
	fmt.Scanln(&N3)

	if N1 < N2 {
		if N2 < N3 {
			fmt.Println("Si")
		}
	} else {
		fmt.Println("No")
	}
}
