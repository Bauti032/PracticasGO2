package main

import "fmt"

func main() {
	var (
		num1 int
		num2 int
	)

	fmt.Print("Ingrese su primer numero: ")
	fmt.Scanln(&num1)
	fmt.Print("Ingrese su segundo numero: ")
	fmt.Scanln(&num2)

	if num1 > num2 {
		fmt.Printf("El numero mayor es: %d. ", num1)
	} else {
		fmt.Printf("El numero mayor es: %d. ", num2)
	}
}
